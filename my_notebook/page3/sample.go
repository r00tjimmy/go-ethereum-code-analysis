package main

import (
  "fmt"
  "time"
  "io"
  "github.com/micro/examples/stream/server/proto"
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

// walk invokes its runTest argument for all subtests in the given directory.
//
// runTest should be a function of type func(t *testing.T, name string, x <TestType>),
// where TestType is the type of the test contained in test files.
// 偶然发现有一个测试的类， 这个类是可以调用某个文件夹， 然后调用文件夹下面的文件， 跑某个测试函数的
func (tm *testMatcher) walk(t *testing.T, dir string, runTest interface{}) {
  // Walk the directory.
  dirinfo, err := os.Stat(dir)
  if os.IsNotExist(err) || !dirinfo.IsDir() {
    fmt.Fprintf(os.Stderr, "can't find test files in %s, did you clone the tests submodule?\n", dir)
    t.Skip("missing test files")
  }
  err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
    name := filepath.ToSlash(strings.TrimPrefix(path, dir+string(filepath.Separator)))
    if info.IsDir() {
      if _, skipload := tm.findSkip(name + "/"); skipload {
        return filepath.SkipDir
      }
      return nil
    }
    if filepath.Ext(path) == ".json" {
      t.Run(name, func(t *testing.T) { tm.runTestFile(t, path, name, runTest) })
    }
    return nil
  })
  if err != nil {
    t.Fatal(err)
  }
}




func Fprintf (w io.writer, format string, args ...interface{}) (int, error)


func Printf(format string, args ...interface{})  (int error) {
  return Fprintf(os.Stdout, format, args...)   
}

func SPrintf(format string, args ...interface{}) string {
  var buf bytes.Buffer
  Fprintf(&buf, format, args...)
  return buf.String()
}


type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
  *c += ByteCounter(len(p))
  return len(p), nil
}

var c ByteCounter
//c.Write([]byte("hello"))
//fmt.Println(c)

//c = 0
//var name = "Dolly"
//fmt.FPrintf(&c, "hello, %s", name)
//fmt.Println(c)


//package io

type Reader interface {
  Read(p []byte) (n int, err error)
}

type Closer interface {
  Close() error
}


type ReadWriter interface {
  Reader
  Writer
}

type ReadWriteCloser interface {
  Reader
  Writer
  Closer
}



type ReadWriter interface {
  Read(p []byte) (n int, err error)
  Write(p []byte) (n int, err error)
}


/**
一个具体的类型可能实现了很多不相关的接口， 考虑在一个组织出售数字文化产品比如音乐， 电影和书籍的程序中可能定义了下列的具体类型：
Album
Book
Movie
Magazine
Podcast
TVEpisode
Track
 */


/**
下面的这些代码等于是定义了一些 接口类型的 struct
 */
type Artifact interface {
  Title() string
  Creators() []string
  Created() time.Time
}

type Text interface {
  Pages() int
  Words() int
  PageSize() int
}

type Audio interface {
  Stream() (io.ReadCloser, error)
  RunningTime() time.Duration
  Format() string
}


type Video interface {
  Stream() (io.ReadCloser, error)
  RunningTime() time.Duration
  Format() string
  Resolution() (x, y int)
}


/**
如果我们发现我们需要以同样的方式来处理 Audio 和 Video， 我们可以定义一个Streamer接口来代表它们之间相同的部分而不必对已存在的类型做改变
 */
type Streamer interface {
  Stream() (io.ReadCloser, error)
  RunningTime() time.Duration
  Format() string
}


type cesiusFlag struct { Celsuis }

func (f *cesiusFlag) Set (s string) error {
  var uint string
  var value falot64
  fmt.Sscanf(s, "%f%s", &value, &uint)

  switch uint {
  case "C", "zero-C":
    f.Celsuis = Celsuis(value)
    return nil

  case "F", "zero-F":
    f.Celsuis = FToC(Fathrenheit(value))
    return nil
  }

  return fmt.Errorf("invalid %q", s)
}


















