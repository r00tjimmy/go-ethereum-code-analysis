package main 

import (
  "fmt"
)

/**
等待多 goroutine 的方案
**/


// 方案1. 只使用单个无缓冲信道阻塞主线

/**

var quit chan int

func foo (id int) {
  fmt.Println(id)
  quit <- 0       // ok, finished
}


func main() {
  coutn := 30 
  quit = make(chan int)     // 无缓冲信道

  // 写入 1000 个值 进信道
  for i := 0; i < coutn; i++ {
    go foo(i)
  }

  // 从信道取出 30 个值
  // for i := 0; i < coutn; i++ {

  //取出 29 个值， 也不会产生死锁
  for i := 0; i < 29; i++ {
    <- quit
  }

**/

  // 为什么上面只是取出了 29 个值， 而不会造成死锁呢？？？
  /**
  是因为 quit 这个信道， 写入的时候， 是用另外的线程（ go foo() )来写的， 而不是用主线程 ( main 线程), 所以说 quit 这个信道，就算不完全取出， 也不会阻塞main进程本身， 因为他们是两个不同的线程， 各自都是独立运行的

  看下面的程序， 假如是同一个main 线程， 这样是会造成死锁的
  **/

// }


/**

func main() {
 var quit chan int = make(chan int)  

 for i := 0; i < 30; i++ {
   <- quit
 } 

 fmt.Println("finished .........")

}

**/

/**

像上面这个程序， 因为是在同一个线程 main 下面的， 所以是会造成死锁的

**/







// 方案2: 把信道换成缓冲 30 的
func main() {
  var quit chan int

  quit = make(chan int, 30)

  for i := 0; i < 30; i++ {     // 因为 quit 是有缓冲的信道， 所以写入之后， 即使不取出， 也不会造成死锁
  // for i := 0; i < 35; i++ {     // 超过了 quit 的信道长度， 产生错误
  // for i := 0; i < 30; i++ {
    quit <- i
    fmt.Println(i)
  }


  for i := 0; i < 20; i++ {
  // for i := 0; i < 40; i++ {
    <-quit
  }

}

/**
所以，会不会造成死锁， 关键看你是什么方式去定义 信道， 有缓冲的信道， 如果你写， 或者读 超过信道的长度， 也是会产生程序错误的

* 无缓冲的信道是一批数据一个一个的「流进流出」
* 缓冲信道则是一个一个存储，然后一起流出去

**/




























