/*
11. Scrivere un programma lez3_quadrati.go che,
dato un intero n, calcola il massimo quadrato <= n
*/

package main

import (
  "fmt"
)

func main() {
  var n int

  fmt.Print("Inserisci n: ")
  fmt.Scan(&n)

  for i:=1;i<=n;i++{
    if i*i > n{
      fmt.Println("Massimo quatrado è:",i-1,":",(i-1)*(i-1))
      break
    }
  }
}
