package main 

import (
  "fmt"
)

const (
  a = iota
  b
  c
)

const (
  d = iota
)


func main() {
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)

  fmt.Println(d)
}
