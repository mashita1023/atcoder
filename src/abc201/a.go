package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
  "reflect"
  "sort"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() string {
  sc.Scan()
  return sc.Text()
}

func typeis(t interface{}) {
  fmt.Println(reflect.TypeOf(t))
}

func atoi(s string) int {
  i, err := strconv.Atoi(strings.TrimSpace(s))
  if err != nil {
    os.Exit(1)
  }
  return i
}

func scanInt() int {
  s := scan()
  return atoi(s)
}

func scanSlice() []string {
  s := scan()
  return strings.Fields(s)
}

func scanSliceInt() []int {
  s := scanSlice()
  var out []int

  for _, v := range s {
    n := atoi(v)
    out = append(out, n)
  }
  return out
}

func removeSlice(s []int, i int) ([]int) {
  return append(s[:i], s[i+1:]...)
}

func main() {
  sl := scanSliceInt()
  
  sort.Slice(sl, func(i, j int) bool { return sl[i] < sl[j]})

  if sl[0] < 1 || sl[2] > 100 {
    os.Exit(1)
  }
  
  if sl[2] - sl[1] == sl[1] - sl[0] {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
  
}
