package main 


import (
  "fmt"
  // "time"
)


func Test() (int, string) {
  return 1, "test_string"
}

func Test2() (int, int) {
  return 1, 1
}


func Test3() (bool, bool) {
  return false, true
}


func MainProcess() {
  a, b := Test()

  fmt.Println(a) 
  fmt.Println(b) 

  if _, c := Test3(); c {
    fmt.Println("xxx") 
  } else {
    fmt.Println("yyy") 
  }
}


func loop(flag string) {
  fmt.Println("this is " + flag)
  for i := 0; i < 10; i++ {
    fmt.Printf("%d\n", i)
  }
}


func main() {
  // case 1
  // loop()
  // loop()


  // case 2
  // go loop()
  // loop()

  // case 3
  // go loop("xx")
  // loop("yy")
  // time.Sleep(time.Second)     // 为什么这里能全部跑完， 是因为循环输出10个数根本不用 一秒


  // case 4
  // 在 python 里面， 等待线程执行完毕的写法如下:
  /**

  for threa ing threads:
    threa.join()
  
  那么我们如果想要实现类似的效果， 也需要一个类似 join 一样额东西来阻塞住主线程， 这个东西就是 channel， 信道
  **/
  case4()
}


func case4() {
  // 建立一个信道
  // var channel chan int = make(chan int)
  // channel ：= make(chan int)

  // 向信道存消息 和 取消息
  var messages chan string = make(chan string)
  go func (message string) {
    messages <- message      // 存消息
  } ("Ping!")

  fmt.Println(<-messages)      // 取消息
}





