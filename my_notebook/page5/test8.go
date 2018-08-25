package main 

import (
  "fmt"
)

/**
缓冲信道 buffered channel， 缓冲信道不仅可以流通数据， 还可以缓存数据， 它是有容量的， 存入一个数据的话， 可以先放在信道里， 不必阻塞当前线程而等待该数据取走

当缓冲信道达到满的状态的时候， 就会出现阻塞， 因为这时再也不能承载更多的数据了， 程序必须要把数据拿走， 才可以流入数据

再声明一个信道的时候， 我们给 make 以第二个参数来指明它的容量， 默认为0， 即无缓冲
**/


// 写入2 个元素都不会阻塞当前 goroutine， 存储个数达到 2 的时候会阻塞
var ch chan int = make(chan int, 2)

/**
func main() {
  ch := make(chan int, 3)
  ch <- 1
  ch <- 2
  ch <- 3

  fmt.Println(<-ch)
  fmt.Println(<-ch)
  fmt.Println(<-ch)
}
**/


// 你会发现上面的代码一个一个的去读信道简直太费事了， go语言允许我们使用range来读取信道

/**
func main() {
  ch := make(chan int, 3)
  ch <- 1 
  ch <- 2 
  ch <- 3

  for v := range ch {
    fmt.Println(v)
  } 
}
**/

/**

如果你执行了上面的代码，会报死锁错误的，原因是range不等到信道关闭是不会结束读取的。也就是如果 缓冲信道干涸了，那么range就会阻塞当前goroutine, 所以死锁咯

**/


/**
func main() {
  ch := make(chan int, 3)
  ch <- 1
  ch <- 2
  ch <- 3
  
  // 下面这个 for 循环其实就有点像去读取 kafka 的topic 一样， 会一直读取，这个在 goroutine  看来， 就是一个死锁， 因为是一直读取的， 所以要避免死锁， 就会判断 信道   的长度，然后合适的逻辑就  break 跳出循环
  for v := range ch {
    fmt.Println(v)
  
    if len(ch) <= 0 {
      break
    }
  }
}
**/

/**
上面的这种方法假如 信道同时 又在读取， 又在写入那就杯具了。。。所以还不是很好的办法
**/



func main() {
  ch := make(chan int, 3)
  ch <- 1
  ch <- 2
  ch <- 3

  // 显式地关闭信道, 这样得话， 信道就不会在等待读取了， 也不能再写入信道了
  close(ch)

  for v := range ch {
    fmt.Println(v)
  }
}

// 被关闭得信道会禁止数据流入， 是只读得， 仍然可以从关闭得信道取出数据， 但是不能再写入数据































