package main 

import (
  "fmt"
)


var complete chan int = make(chan int)

func loop() {
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }

  complete <- 0     // 执行完毕， 发了个消息

  // 上面这句发了一个信息给信道， 如果想阻塞住这个 loop 函数的话， 那么就应该来这个信道拿这个信息， 如果不拿这个信息的话， 是阻塞不了 loop 函数的
}


func main() {
  go loop()
  <- complete   // 直到线程跑完， 取到消息， main在此组塞住 

  // go loop() 是跑在另外的线程上的， 假如想取得 go loop() 的运行逻辑， 那么就必须去  go loop() 这个线程的信道里去， 去 取得信道里的数据（取得信道里的数据， 其实本质的作用就是要 阻塞这个信道）， 从而能获取 go loop()  的执行结果
}