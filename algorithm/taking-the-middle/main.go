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

func main() {
  
  n, err := IntStdin()

  if err != nil {
    fmt.Println(err)
    return
  } else if len(strconv.Itoa(n)) > 5 || n < 1{
    fmt.Println("dame")
    return
  }

  arr, err := IntArrayStdin()
  
  if err != nil {
    return
  }
  
  fmt.Println(len(arr))
  if len(arr) != (2*n) {
    fmt.Println("please tyr again.")
    return
  }
  
  var ans int = 0

  for i := 0; i < n; i++ {
    idx := MaxArrayNum(arr)
    ans += arr[idx]
    arr = RemoveSlice(arr, idx)

    idx = len(arr) / 2
    arr = RemoveSlice(arr, idx)

  }

  fmt.Printf("%v\n",ans)
}
