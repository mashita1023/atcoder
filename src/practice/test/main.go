package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func scan() string {
  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()

  return stdin.Text()
}

func spaceSlice(s string) []string {
  slice := strings.Fields(s)
  return slice
}

func atoi(s string) int {
  i, err := strconv.Atoi(s)
  if err != nil {
    return 0
  }
  return i
}

func main() {
  s := scan()
  first := atoi(s)
  slice := scan()
  sl := spaceSlice(slice)
  second := atoi(sl[0])
  third := atoi(sl[1])
  ans := first + second + third

  //str := scan()
  var str string
  fmt.Scanf("%s", &str)
  fmt.Printf("%v %v",ans,str)
}
