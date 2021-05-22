package main
import (
  "fmt"
  "os"
  "bufio"
  "reflect"
  "strconv"
  "strings"
)

func P(t interface{}) {
  fmt.Println(reflect.TypeOf(t))
}

func StrStdin() (res string){
  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  text := stdin.Text()
  return text
}

func IntLength(n int) (int){
  return len(strconv.Itoa(n))
}

func IntStdin() (int, error) {
  input := StrStdin()
  return strconv.Atoi(strings.TrimSpace(input))
}

func main() {
  num, err := IntStdin()

  if err != nil {
    fmt.Println(err)
    return
  } else {
    P(num)
  }

  var l int = IntLength(num)
  
  if l > 15 || num < 1 {
    fmt.Println("this number is not 1 <= N <= 10^15. plz try again!")
    return
  }

  var ans int = 0
  
  if l < 4 {
    fmt.Println(ans)
    return
  }
  ans += num - 999
  
  if l < 7 {
    fmt.Println(ans)
    return
  }
  ans += (num - 999999)
  
  if l < 10 {
    fmt.Println(ans)
    return
  }
  ans += (num - 999999999)
  
  if l < 13 {
    fmt.Println(ans)
    return
  }
  ans += (num - 999999999999)

  fmt.Println(ans)
}
