
/**

type有如下几种用法：
1. 定义结构体
2. 定义接口
3. 类型定义
4. 类型别名
5. 类型查询

**/


// 定义结构体
type name struct {
  field1    dataType
  field2    dataType
  field3    dataType 
}



// 定义接口
type name interface {
  Read()
  Write()
} 


// 定义类型
type name string

// 类型定义可以在原类型的基础上创造出新的类型，有些场合下可以使代码更加简洁，如下边示例代码：

type handle func(str string)

func exec(f handle) {
  f("hello")
}


func main() {
  var p = func(str string) {
    fmt.Println("first", str)
  }
  exec(p)


  exec(func(str string) {
    fmt.Println("second", str)
  })
}

// 看下面的代码， 因为重新定义了类型，而使代码简洁了
type handle func(str string, str2 string, num int, money float64, flag bool)


func exec(f handle) {
  f("hello", "world", 10, 11.23, true)
}


// 下面是具体 func 的实现
func demo(str string, str2 string, num int, money float64, flag bool) {
  fmt.Println(str, str2, num, money, flag)
}

func main() {
  exec(demo)
}




// 类型查询
func main() {
  var a interface{} = "abc"

  switch v := a(type) {
  case string:
    fmt.Println("字符串")
  case int:
    fmt.Println("整型")
  default:
    fmt.Println("其他类型", v)

  }
}



  









