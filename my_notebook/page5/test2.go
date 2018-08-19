package main 

import (
  "fmt"
)


var ch chan int = make(chan int)

func  foo() {
  ch <- 0   // 向 ch 中加数据， 如果没有其他 goroutine 来取走这个数据， 那么挂起foo， 直到main函数把0这个数据拿走 
}


func main() {
  go foo()
  
  
  // 这里的这个写法即表示 主线程要等待 ch 这个信道读取完成， 因为 ch 这个信道的  写 数据是在函数 foo 里面的， 意思就是表示要等待 foo 函数执行完成。
  <- ch    // 从ch取走数据， 如果ch中还没放数据， 那就挂起 main 线， 直到 foo 函数中放数据为止,
  fmt.Println("xx")
}