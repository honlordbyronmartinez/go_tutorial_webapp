package main
//git test comment
import "fmt"

func foo(c chan int, somevalue int){
  c <- somevalue * 5
}

func main ()  {
  fooVal := make(chan int)

  go foo(fooVal, 5)
  go foo(fooVal, 3)

  v1 := <- fooVal
  v2 := <- fooVal

  fmt.Println( v1, v2)
}
