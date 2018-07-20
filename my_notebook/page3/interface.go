package main

import (
  "fmt"
)


/**



// 1
type I interface {
  Get() int
  Set(int)
}


// 2
type S struct {
  Age int
}


func (s S) Get() int {
  return s.Age
}


// 3
func (s *S) Set(age int) {
  s.Age = age
}


func f(i I) {
  i.Set(10)
  fmt.Println(i.Get())
}


**/



/**
这段代码在 #1 定义了 interface I，在 #2 用 struct S 实现了 I 定义的两个方法，接着在 #3 定义了一个函数 f 参数类型是 I，S 实现了 I 的两个方法就说 S 是 I 的实现者，执行 f(&s) 就完了一次 interface 类型的使用。
**/


/**

func main() {
  s := S{}    //不用显式定义s的类型

  // 也可以显式定义s的类型
  var i I 
  i = &s

  f(&s) 
}



func test() {
  var data_slice []int = foo()
  //定义好  interface 数组的长度 就可以跟 slice 切换了
  var interface_slice []interface{} = make([]interface, len(data_slice))

  for i, d := range data_slice {
    interface_slice[i] = d
  }
}



switch t := i.(type) {
case *S:
  fmt.Println("i store *S", t)
case *R:
  fmt.Println("i store *R", t)
}

**/




type I interface {
  Get() int
  Set(int)
}

type SS struct {
  Age int
}

func (s SS) Get() int {
  return s.Age
}

func (s SS) Set(age int) {
  s.Age = age
}

func f(i I) {
  i.Set(5)
  fmt.Println(i.Get())
}

func main(){
  ss := SS{}
  f(&ss) //ponter
  f(ss)  //value
}
















