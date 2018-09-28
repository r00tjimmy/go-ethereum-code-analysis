package main 

import (
  "fmt"
  "time"
  "golang.org/x/net/context"
)


// 模拟一个最小执行时间的堵塞函数
func inc(a int) int {
  res := a + 1
  time.Sleep(1 * time.Second)
  return res
}


func Add(ctx context.Context, a, b int) int {
  res := 0
  for i := 0; i < a; i++ {
    res = inc(res) 

    select {
    case <- ctx.Done():
      return -1
    default:
    }
  }

  for i := 0; i < b; i++ {
    res = inc(res)

    select {
    case <- ctx.Done():
      return -2
    default:
    }
  }

  return res
}


func main() {
  // {
  //   // 使用 golang开放的API 计算 a+b
  //   a := 1
  //   b := 2
  //   // 设置 timeout 的时间
  //   timeout := 6 * time.Second
  //   ctx, _ := context.WithTimeout(context.Background(), timeout)
  //   /**
  //   检查哪个线程等待的时间比较久
  //   **/
  //   // res := Add(ctx, 1, 2)
  //   res := Add(ctx, 4, 2)
  //   fmt.Printf("compute: %d + %d, result: %d\n", a, b, res)
  // }


  {
    // 手动取消
    a := 1
    b := 2
    ctx, cancel := context.WithCancel(context.Background())

    go func() {
      time.Sleep(1 * time.Second)
      cancel()
    }()

    res := Add(ctx, 1, 2)
    fmt.Printf("compute xxxx: %d + %d, result:  %d\n", a, b, res)
  }


}










