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

func permutation(n int, k int) int {
  v := 1
  if 0 < k && k <= n {
    for i := 0; i < k; i++ {
      v *= n - i
    }
  } else if k > n {
    v = 0
  }
  return v
}

func factorial(n int) int {
  return permutation(n, n - 1)
}

func main() {
  sl := scanSliceInt()

  var p int = factorial(sl[0] + sl[1])
  fmt.Println(sl)
  fmt.Println(p)
}
