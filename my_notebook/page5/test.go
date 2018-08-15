package main 


import (
  "fmt"
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


func main() {
  a, b := Test()

  fmt.Println(a) 
  fmt.Println(b) 

  if _, c := Test3(); c {
    fmt.Println("xxx") 
  } else {
    fmt.Println("yyy") 
  }
}