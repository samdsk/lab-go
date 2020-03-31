/*
Caratteri ASCII
*/


package main

import (
    "fmt"
)

func main() {
  var carattere int32
  carattere = 65
  fmt.Println(string(carattere))
  carattere++
  fmt.Println(string(carattere))  
}
