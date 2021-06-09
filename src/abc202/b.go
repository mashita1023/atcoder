package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() string {
  sc.Scan()
  return sc.Text()
}

func convert(str string) string {
  if str == "6" {
    return "9"
  } else if str == "9" {
    return "6"
  }
  return str
}

func main() {
  inp := scan()

  str := strings.Split(inp, "")

  lnt := len(str)

  var out []string = make([]string, lnt, lnt)
  
  for i := 0; i < lnt ; i++ {
    out[i] = convert(str[lnt - i - 1])
  }

  fmt.Println(strings.Join(out, ""))
  
}
