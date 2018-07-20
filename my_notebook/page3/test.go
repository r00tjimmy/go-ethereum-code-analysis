package main

import (
  "fmt"
)


type I interface {
  Get() int
  Set(int)
}



type S struct {
  Age int
}


func (s S) Get() int {
  return s.Age
}



func (s *S) Set(age int) {
  s.Age = age
}


func f(i I) {
  i.Set(10)
  fmt.Println(i.Get())
}


func main() {
  s := S{}
  f(&s) 
}
