package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
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
  return strconv.Atoi(strings.TrimSpace(s))
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

  for _, v := r s {
    n, _ := atoi(v)
    out = append(out, n)
  }
  return out
}

func removeSlice(s []int, i int) []int {
  return append(s[:i], s[s+1:]...)
}

func main() {
  
}
