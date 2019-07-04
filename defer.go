package main
// test
import ("fmt"
        "time")

func foo(){
  for i := 0; i < 3 ; i++{
    time.Sleep(time.Millisecond*100)
    fmt.Println(i)
  }
}

func main(){
  foo()
}
