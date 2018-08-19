package main 

import (
  "fmt"
)


var ch1 chan int = make(chan int)
var ch2 chan int = make(chan int)

func say(s string) {
  fmt.Println(s)
  ch1 <- <- ch2       // ch1 等待 ch2 流出的数据  
}


func main() {
  go say("hello") 
  <- ch1        // 这里肯定死锁， 因为 ch2 是空的， 本来不会流出数据给 ch1， 所以 ch1 也就没有写入

  // 其中主线等ch1中的数据流出，ch1等ch2的数据流出，但是ch2等待数据流入，两个goroutine都在等，也就是死锁。

}