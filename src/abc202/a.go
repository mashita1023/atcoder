package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
  "reflect"
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

func dice(i int) int {
  return 7 - i
}

func main() {
  slice := scanSliceInt()

  a := dice(slice[0])
  b := dice(slice[1])
  c := dice(slice[2])

  fmt.Println(a + b + c)
}
