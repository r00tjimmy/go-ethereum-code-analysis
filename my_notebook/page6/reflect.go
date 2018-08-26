package main 

import (
  "fmt"
  "reflect"
)


/**

func WhatIsYourType(i interface{}) string {
  val := reflect.ValueOf(i)
  return val.Kind().String()
}



func main() {
  str := []string{"go", "lang", "reflect", "sample"}

  // val := reflect.ValueOf(str)
  // fmt.Println(val)

  it := 2009
  sit := map[string]string{"hi": "girl", "hello": "sir"}

  fmt.Println(WhatIsYourType(str))
  fmt.Println(WhatIsYourType(it))
  fmt.Println(WhatIsYourType(sit))
}

**/

type Ref struct {
  key string
  val string
}


func (this *Ref) SayHi(str interface{}) {
  fmt.Println("reflect say", str)
}


func main() {
  ref := &Ref{key: "hi", val: "girl"}     // ref 是一个指向结构体的指针

  iref := reflect.ValueOf(ref)
  fn := iref.MethodByName("SayHi")

  var str = "hello reflect"
  var in []reflect.Value
  in = append(in, reflect.ValueOf(str))

  // 再增加这里的逻辑， 会造成错误, 传递给函数的参数太多了
  // var str2 = "my test"
  // in = append(in, reflect.ValueOf(str2))


  fn.Call(in)
}












