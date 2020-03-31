package main

import (
    "fmt"
)

func main() {
  n := 10
  n2 := 2*n

  for i:=n;i>0;i--{
    for j:=0;j<=n2;j++{
      if j > n2-i+1{break}

      if j == i || (j == i+1 && i != n){
        fmt.Print("/")
      }else if j == n2-i+1 || j == n2-i{
        fmt.Print("\\")
      }else{
        fmt.Print(" ")
      }
    }
    fmt.Println()
  }
}
