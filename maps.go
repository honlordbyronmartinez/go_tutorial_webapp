package main

import ("fmt")

func main() {
  grades := make(map[string]float32)

  grades["Yolanda"] = 42
  grades["Byron"] = 40
  grades["Emeni"] = 12
  grades["Amyah"] = 10

  for k, v := range grades{
    fmt.Println(k,":",v)
  }
}
