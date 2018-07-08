package main

import (
  "fmt"
)

func GetStrs() (string, string, string) {
  // 返回多个字符串的用法
  return "p", "k", "m"
}


func GetStrsSli() []string {
  // 返回多个字符串的用法
  return []string{"p", "k", "m"}
}


func main() {
  person_salary := map[string]int {
    "steve": 120000,
    "jamie": 150000,
  }

  person_salary["mike"] = 40000
  fmt.Println(person_salary)

  new_person_salary := person_salary
  new_person_salary["mike"] = 80000
  fmt.Println(person_salary)
  fmt.Println(new_person_salary)


  sli := []string{"a", "b", "c"}
  sli = append(sli, "d", "e")
  fmt.Println(sli)

  s := "f, g, h"
  sli = append(sli, s)
  fmt.Println(sli)

  fmt.Println(len(sli))

  sli = append(sli, GetStrsSli()...)
  // ...变量， ...参数的用法
  fmt.Println(sli)
  fmt.Println(len(sli))
}


