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

type Mountain struct {
  name string
  height int
}

func main() {

  n := scanInt()

  arr := make([]Mountain, 0)

  for i := 0; i < n; i++ {
    sl := scanSlice()

    var m Mountain
    m.name = sl[0]
    m.height = atoi(sl[1])

    arr = append(arr, m)
  }

  sort.Slice(arr, func(i, j int) bool {return arr[i].height > arr[j].height})

  fmt.Println(arr[1].name)
  
}
