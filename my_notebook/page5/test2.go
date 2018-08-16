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
 
 <- ch    // 从ch取走数据， 如果ch中还没放数据， 那就挂起 main 线， 直到 foo 函数中放数据为止

 fmt.Println("xx")
}