package main

import (
  "fmt"
  "os"
  "palyn/gen"
  "strconv"
)

func main() {
  n,_ := strconv.Atoi(os.Args[1])
  fmt.Println(gen.IterPalyn(n))
  fmt.Println(gen.ResPalyn(n))
}
