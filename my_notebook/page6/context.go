
// context包

/**
一般用于在不同的 channel 之间传递数据的
**/

/**
type Context interface {
  // 返回已经取消了的
  Done() <- chan struct{}

  Err() error

  Deadline() (deadline time.Time, ok bool)

  Value(key interface{}) interface{}
}


func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
  c := newCancelCtx(parent)
  propagateCancel(parent, &c)
  return &c, func() { c.cancel(true, Canceled) }
}


func newCancelCtx(parent Context) cancelCtx {
  return cancelCtx {
    Context:    parent,
    done:       make(chan struct{}), 
  }
}

**/



// type Context interface {
//   Deadline() (deadline time.Time, ok bool) 
//   Done() <- chan struct{},
//   Err() error,
//   Value(key interface{}) interface{}
// }


package main 

import (
  "fmt"
  "time"
  "golang.org/x/net/context"
)


func main() {
  // context.Background 为 上下文的 根节点 
  // 定义了 ctx cancal的属性
  ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(time.Second * 5)) 

  // 定义 ctx 的上下文传递的值
  ctx = context.WithValue(ctx, "Test", "123456")

  if t, ok := ctx.Deadline(); ok {
    fmt.Println(time.Now())
    fmt.Println(t.String())
  }

  // 在当前的 groutine 输出上下文传递的值
  go func(ctx context.Context) {
    fmt.Println(ctx.Value("Test"))
    fmt.Println("waitting for this context groutine finished --------------- [main]")

    for {
      select {
      case <- ctx.Done():
        fmt.Println("this groutine done......")
        fmt.Println(ctx.Err())
        return
      }
    }

  } (ctx)

  time.Sleep(time.Second * 5)     // sleep 3 seconds
  fmt.Println("finished!!!")
  cancelFunc()

}



























