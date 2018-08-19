/**
关于死锁的一些代码
**/
package main

import (
  "fmt"
)

/**
func main() {
  ch := make(chan int)
  <- ch   // 阻塞main groutint， 信道ch被锁,  因为 ch 本来就是空的， 只是初始化好了，但是没有东西， 如果读取不到东西的话， 那么就会造成死锁

  // 在单一的 goroutine 里面 只写不读 channel， 或者只读不写 channel， 两种情况都会造成死锁

  fmt.Println("done ............. ")
}
**/


func main() {
  ch := make(chan int)
  ch <- 1             // 这种就是只写不读的情况， 也会造成死锁
  fmt.Println("done ............ ")
}