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


