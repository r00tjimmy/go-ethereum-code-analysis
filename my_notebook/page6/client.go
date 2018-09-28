
// RPC客户端源码

type Client struct {
  idCounter         uint32          // 属于这个客户端的id标识

  // 生成链接的函数， 客户端会调用这个函数来生成一个网络连接对象
  connectFunc       func(ctx context.Context) (net.Conn, error)
  // HTTP协议和非HTTP协议有不同的处理流程， HTTP不支持长连接， 只支持一个请求对应一个回应的模式， 同时也不支持发布/订阅模式
  isHTTP           bool


  //通过这里的注释可以看到，writeConn是调用这用来写入请求的网络连接对象，
  //只有在dispatch方法外面调用才是安全的，而且需要通过给requestOp队列发送请求来获取锁，
  //获取锁之后就可以把请求写入网络，写入完成后发送请求给sendDone队列来释放锁，供其它的请求使用。
  writeConn         net.Conn


  // for dispatch
  //下面有很多的channel，channel一般来说是goroutine之间用来通信的通道，后续会随着代码介绍channel是如何使用的。
  close       chan struct{}
  didQuit     chan struct{}                  // closed when client quits
  reconnected chan net.Conn                  // where write/reconnect sends the new connection
  readErr     chan error                     // errors from read
  readResp    chan []*jsonrpcMessage         // valid messages from read
  requestOp   chan *requestOp                // for registering response IDs
  sendDone    chan error                     // signals write completion, releases write lock
  respWait    map[string]*requestOp          // active requests
  subs        map[string]*ClientSubscription // active subscriptions

}


func newClient(initctx context.Context, connectFunc func(context.Context) (net.Conn, error))  (*Client, error) {

  conn, err := connectFunc(initctx) 
  if err != nil {
    return nil, err
  }

  _, isHTTP := conn.(*httpConn)

  c := &Client {
    writeConn:          conn,
    isHTTP:             isHTTP,
    connectFunc:        connectFunc,
    close:              make(chan struct{}),
    didQuit:            make(chan struct{}),
    reconnected:        make(chan net.Conn),
    readErr:            make(chan error),
    readResp:           make(chan []*jsonrpcMessage),
    requestOp:          make(chan *requestOp),
    sendDone:           make(chan error, 1),
    respWait:           make(map[string]*requestOp),
    subs:        make(map[string]*ClientSubscription),
  }

  // 如果不是 HTTP 连接， 则进行 dispatch 分发
  is !isHTTP {
    go c.dispatch(conn)
  }

  return c, nil

}


func (c *Client) Call (result interface{}, method string, args ...interface{}) error {
  // 上下文的根节点
  ctx := context.Background()
  return c.CallContext(ctx, result, method, args...)  
}


func (c *Client) CallContext (ctx context.Context, result interface{}, method string, args ...interface{}) error {
  msg, err := c.newMessage(method, args...)

  if err != nil {
    return err
  }

  op := &requestOp{ ids: []json.RawMessage{msg.ID}, resp: make(chan *jsonrpcMessage, 1) }

  if c.isHTTP {
    err = c.sendHTTP(ctx, op, msg)
  } else {
    err = c.send(ctx, op, msg)
  }

  if err != nil {
    return err
  }

  // dispatch has accepted the request and will close the channel it when it quit
  switch resp, err := op.wait(ctx); {
  case err != nil :
    return err
  case resp.Error != nil: 
    return resp.Error
  case len(resp.Result) == 0:
    return ErrNoResult
  default:
    return json.Unmarshal(resp.Result, &result)
   
  }

}


// sendHTTP,这个方法直接调用doRequest方法进行请求拿到回应。然后写入到resp队列就返回了。
func (c *Client) sendHTTP(ctx context.Context, op *requestOp, msg interface{}) error {
  hc := c.writeConn.(*httpConn)
  respBody, err := hc.doRequest(ctx, msg)
  if err != nil {
    return err
  }
  defer respBody.Close()
  var respmsg jsonrpcMessage
  if err := json.NewDecoder(respBody).Decode(&respmsg); err != nil {
    return err
  }
  op.resp <- &respmsg
  return nil
}



func (op *requestOp) wait(ctx context.Context) (*jsonrpcMessage, error) {
  select {
  case <- ctx.Done():
    return nil, ctx.Err()
  case resp := <- op.resp:
    return resp, op.err
  }
}











