


type CalculatorService struct {

}



func (s *CalculatorService) Add (a, b int) int {
  return a + b
}



func (s *CalculatorService) Div(a, b int) (int, error) {
  if b == 0 {
    return 0, errors.New("divide by zero")
  } 
  return a / b, nil
}



calculator := new(CalculatorService)      // 创建一个新的RPC具体执行对象
server := NewServer()                     // 创建一个新的服务对象 
server.RegisterName("calculator", calculator)       // 把RPC注册到服务对象去， 以便RPC可以调用


// 监听unix 本地域
l, _ := net.ListenUnix("unix", &net.UnixAddr{Net: "unix", Name: "/tmp/calculator.sock"})
for {
  /**
  下面这里就很秒， 把socket连接， json格式， RPC服务 三个逻辑都封装开了， 耦合度比较低
  **/
  c, _ := l.AcceptUnix()
  codec := v2.NewJSONCodec(c)
  go server.ServeCodec(codec)         // 写入并行的信道
}


/**

订阅在下面几种情况下会被删除

用户发送了一个取消订阅的请求
创建订阅的连接被关闭。这种情况可能由客户端或者服务器触发。 服务器在写入出错或者是通知队列长度太大的时候会选择关闭连接。


**/

