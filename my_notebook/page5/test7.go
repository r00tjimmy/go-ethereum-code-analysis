package main 

import (
  "fmt"
)


var ch  chan int = make(chan int)


func foo(id int) {
  ch <- id
}


func main() {
  // 写入信道
  for i := 0; i < 5; i++ {
    go foo(i)
  } 

  // 输出信道
  for i := 0; i < 5; i++ {
    // 从信道取处理的数据是无序的， 可见其实信道底层为了性能考虑， 应该也是采用了散列哈希来实现的
    fmt.Println(<- ch)
  }
}

//我们开了5个goroutine，然后又依次取数据。其实整个的执行过程细分的话，5个线的数据 依次流过信道ch, main打印之, 而宏观上我们看到的即 无缓冲信道的数据是先到先出，但是 无缓冲信道并不存储数据，只负责数据的流通















