package main

import (
  "fmt"
  "os"
)

func main() {
  j := len(os.Args)
  for i := 1; i < j; i ++ {
    arg := os.Args[i]
    fmt.Println(arg)
  }
}
