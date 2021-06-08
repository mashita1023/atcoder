package main
import (
  "fmt"
  "os"
  "bufio"
  "reflect"
  "strconv"
  "strings"
  "errors"
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

func ArrayStdin() ([]string) {
  var input string = StrStdin()
  var slice []string = strings.Fields(input)
  return slice
}

func IntArrayStdin() ([]int, error) {
  input := ArrayStdin()
  var output []int
  for _, v := range input {
    n, _ := strconv.Atoi(v)

    if len(v) > 10 || n < 0 {
      return output, errors.New("try again.")
    }

    output = append(output, n)
  }

  return output, nil
}

func MaxArrayNum(arr []int) (int) {
  var max int = 0
  var idx int
  
  for i, v := range arr {
    if max < v {
      max = v
      idx = i
    }
  }

  return idx
}

func RemoveSlice(slice []int, s int) ([]int) {
  return append(slice[:s], slice[s+1:]...)
}

func Gcd(a int, b int) int {

  var s int = 0

  for {
    s = b % a
    //    fmt.Printf("a: %v, b:%v, s:%v\n", a, b, s)
    if s == 0 {
      return a
    }
    
    b = a
    a = s
    
  }
  
}

func main() {

  arr, err := IntArrayStdin()

  if err != nil {
    return
  }

  if arr[0] < 1 || arr[0] > arr[1] || arr[1] > 200000 {
    fmt.Println("dame")
    return
  }
  var max int = 1
  var a int
  for i:=arr[0]; i<arr[1]-1; i++ {
    for j:=arr[0]+1; j<=arr[1]; j++ {
      //      fmt.Println("gcd")
      a = Gcd(i, j)
      if i == j{
        continue
      } else if max < a {
        max = a
        fmt.Printf("%v : %v => max: %v\n", i, j, max)
      }
    }
  }

  fmt.Println(max)
}
