




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