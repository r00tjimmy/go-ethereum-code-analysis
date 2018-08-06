/**
golang 的位运算范例

&      位运算 AND
|      位运算 OR
^      位运算 XOR
&^     位清空 (AND NOT)
<<     左移
>>     右移

**/

package main 

import (
  "fmt"
)


func testFunc1() {
  // x&^y==x&(^y) 首先我们先换算成2进制  0000 0010 &^ 0000 0100 = 0000 0010 如果ybit位上的数是0则取x上对应位置的值， 如果ybit位上为1则结果位上取0
  x := 2
  y := 4
  fmt.Println(x&^y)
}


func testFunc2() {
  // 进行转化为二进制 然后向左或者向右移动。
  x := 2
  y := 4

  fmt.Println(x << 1)
  fmt.Println(y >> 1)
}


func main() {
  x := 4
  fmt.Println(^x)     //二进制反转

  y := 4
  z := 2
  // XOR是不进位加法计算，也就是异或计算。0000 0100 + 0000 0010 = 0000 0110 = 6
  fmt.Println(y^z)

}


