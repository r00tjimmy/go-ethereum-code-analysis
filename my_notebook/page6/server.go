




type serviceRegistry map[string] *service      // collection of services

// 这个是有客户端的RPC请求过来， 就会从这里去取得调用的逻辑
type callbacks map[string] *callback         // collection of RPC callbacks


// 这个是表示， 有什么客户端订阅了RPC， 就从这里去取得调用的逻辑
type subscriptions map[string] *callback 

// RPC server 的核心数据结构
type Server struct {
  services serviceRegistry

  run int32
  codecsMu sync.Mutex       // 并行逻辑的锁
  codecs *set.Set
}


// callback is a method callback which was registered in the server
type callback struct {
  rcvr              reflect.Value
  method            reflect.Method
  argTypes          []reflect.Type
  hasCtx            bool 
  errPos            int
  isSubscribe       bool 
}


// service represents a registered object
type service struct {
  name            string
  typ             reflect.Type
  callbacks       callbacks
  subscriptions   subscriptions
}


const MetadataApi = "rpc"
// NewServer will create a new server instance with no registered handles
func NewServer() *Server {
  server := &Server {
    services:         make(serviceRegistry),      // 实际上这个就是一个 string 的map， 储存RPC服务的名称的
    codecs:           set.New(),
    run:              1,
  }

  rpcService := &RPCService(server)
  server.RegisterName(MetadataApi, rpcService)

  return server
}




func (s *Server) RegisterName(name string, rcvr interface{}) error {
  if s.services == nil {
    s.services = make(serviceRegistry)
  }

  svc := new(service)
  svc.typ = reflect.TypeOf(rcvr)
  rcvrVal := reflect.ValueOf(rcvr)

  if name == "" {
    return fmt.Errorf("no serivce error ----------------------- %s", svc.typ.String())
  }


  # reflect.Indirect 是通过反射的形式获得变量的指针
  if !isExported(reflect.Indirect(rcvrVal).Type().Name()) {
    return fmt.Errorf("%s is not exported ", reflect.Indirect(rcvrVal).Type().Name())
  }

  methods, subscriptions := suitableCallbacks(rcvrVal, svc.typ)

  if regsvc, present := s.services[name]; present {
    // golang 几乎有一个特点就是，每个逻辑开始之前，都要做一下致命的错误判断？
    if len(methods) == 0 && len(subscriptions) == 0 {
      return fmt.Errorf("services dont have method %T", rcvr)
    }

    for _, m := range methods {
      regsvc.callbacks[formatName(m.method.Name)] = m 
    }

    for _, s := range subscriptions {
      regsvc.subscriptions[formatName(s.method.Name)] = s
    }

    return nil
  }

  svc.name = name
  svc.callbacks, svc.subscriptions = methods, subscriptions

  if len(svc.callbacks)  == 0 && len(svc.subscriptions) == 0 {
    return fmt.Errorf("service T doesnot ", rcvr)
  }

  s.services[svc.name] = svc
  return nil

}



// 一般回调的做法都是通过反射来获取一些信息， 例如这里就是通过反射来获取合适的方法, 下面的 suitableCallbacks 就是这样的一个逻辑

func suitableCallbacks(rcvr reflect.Value, typ reflect.Type) (callbacks, subscriptions) {
  // 这个是回调的
  callbacks := make(callbacks)

  // 这个是订阅的
  subscriptions := make(subscriptions)


  // golang 的这种写法类似 C++ 里面的 goto
  METHODS:
  // 遍历方法？？？
  for m := 0; m < typ.NumMethod(); m++ {
    method := typ.Method(m)
    mtype:= method.Type
    mname := formatName(method.Name)
    if method.PkgPath != "" {
      continue
    }

    var h callback
    h.isSubscribe = isPubSub(mtype)
    h.rcvr = rcvr
    h.method = method
    h.errPos = -1

    firstArg := 1
    numln := mtype.Numln()
    if numln >= 2 && mtype.ln(1) == contextType {
      h.hasCtx = true
      firstArg = 2      
    }

    if h.isSubscribe {
      h.argTypes = make([]reflect.Type, numln - firstArg)
      for i := firstArg; i < numln; i++ {
        argType := mtype.ln(i)
        if isExportedOrBuildtinType(argType) {
          h.argTypes[i - firstArg] = argType
        } else {
          continue METHODS
        }
      }

      subscriptions[mname] = &h
      continue METHODS
    }

  }

}


// 核心的代码逻辑， serverRequest 方法， sync.WaitGroup实现了一个信号量的功能， Context实现上下文管理
func (s *Server) serveRequest(codec ServerCodec, singleShot bool, options CodecOption) error {
  var pend sync.WaitGroup 

  defer func() {
    if err := recover(); err != nil {
      const size = 64 << 10         # 左移 10 位??
      buf := make([]byte, size)
      buf = buf[:runtime.Stack(buf, false)]
      log.Error(string(buf))
    }

    s.codecsMu.Lock()
    s.codecs.Remove(codecs)
    s.codecsMu.Unlock()
  }()

  ctx, cancel := context.WithCancel(context.Backgroud())
  defer cancel()

  if options & OptionSubscriptions == OptionSubscriptions {
    // context 是获取不同的 goroutine 的值
    ctx = context.WithValue(ctx, notifierKey{}, newNotifier(codec)) 
  }

  s.codecsMu.Lock()
  if atomic.LoadInt32(&s.run) != 1 {      // server stopped
    s.codecsMu.Unlock()
    return &shutdownError{}
  }
  s.codecs.Add(codec)
  s.codecsMu.Unlock()

  for atomic.LoadInt32(&s.run) == 1 {
    reqs, batch, err := s.readRequest(codec)
    if err != nil {
      if err.Error() != "EOF" {
        log.Debug(fmt.Sprintf("read error %v/n", err))
        // codec 只是一个序列， 是一个回复, 是RPC的一个信息加工过之后的回复
        codec.Write(codec.CreateErrorResponse(nil, err))
      }
      // 这里主要是考虑多线程处理的时候等待所有的 rquest处理完毕
      // 每启动一个go线程会调用pend.Add()
      // 处理完成后调用 pend.Done() 会减1， 当为0的时候， Wait()方法就会返回
      pend.Wait()
      return nil
    }
  }


  // 启动线程对请求进行服务
  go func(reqs []*serveRequest, batch bool) {
    defer pend.Done()
    if batch {
      s.execBatch(ctx, codec, reqs)
    } else {
      s.exec(ctx, codec, reqs[0])
    }
  }(reqs, batch)

}


// 定义返回的数据类型
type rpcRequest struct {
  service       string
  method        string
  id            interface{}
  isPubSub      bool
  params        interface{}
  err           Error
}




type serverRequest struct {
  id                interface{}
  svcname           string
  callb             *callback
  args              []reflect.Value
  isUnsubsribe      bool
  err               Error 
}



func (s *Server) readRequest(codec ServerCodec) ([] *serveRequest, bool, Error) {
  
}






















