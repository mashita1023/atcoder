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
  n := scanInt()
  a := scanSliceInt()
  b := scanSliceInt()
  c := scanSliceInt()
  var ans int
  
  for _, v := range a {
    for _, w := range c {
      if v == b[w - 1] && v <= n {
        ans++
      }
    }
  }

  fmt.Println(ans)
}
