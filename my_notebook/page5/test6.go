package main 

import (
  "fmt"
)


c, quit := make(chan int), make(chan int)

go func() {
  c <-          // c 通道的数据没有被其他groutine读取走， 阻塞当前goroutine 
  quit <- 0     // quit 因为上面的错误逻辑， quit通信始终没有办法写入数据 
}()


<- quit   // quit 等待数据的写
<- c       // 想要避免死锁， 还需要再加上这一句， 取走 c 信道的数据就可以了


/**
仔细分析的话，是由于：主线等待quit信道的数据流出，quit等待数据写入，而func被c通道堵塞，所有goroutine都在等，所以死锁。
简单来看的话，一共两个线，func线中流入c通道的数据并没有在main线中流出，肯定死锁。
**/


// 但是， 是否所有不成对的 信道存取数据的情况 都会造成死锁呢？？？

func main() {
  c := make(chan int)

  go func() {
    c <- 1
  }()
}

// 上面这个不会造成死锁， 因为main没有等待其他的 groutine， 自己先跑完了， 所以没有数据流入c通道




// 假如我们设置一个缓冲信道， 也就是说， 放入一个数据， c并不会挂起当前线， 因为信道本身就可以储存一个数据
c := make(chan int, 1)

//  但是假如这种情况下， 再往 c 信道写入多一个数据的话， 那么就一定要再读取出来， 不然也会造成阻塞， 因为 c 信道最多只能缓冲（储存）一个数据



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
    fmt.Println(<- ch)
  }
}






















