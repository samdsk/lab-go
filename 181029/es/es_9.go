/*
9. Scrivere un programma lez3_fibonacci.go che, dato un intero n,
stampa i primi n numeri della serie di Fibonacci: n_i = n_{i-1} + n_{i-2}
*/

package main

import (
    "fmt"
)

func main() {
  var n_1,n_2,n int64 = 1,1,0

  fmt.Print("Inserisci un numero > 2: ")
  fmt.Scan(&n)
  fmt.Println("Serie di Fibonacci fino",n)

  fmt.Println("1 : 1")
  fmt.Println("2 : 1")

  for i:=2;i<int(n);i++{
    var temp int64 = n_1+n_2
    n_1 = n_2
    n_2 = temp

    fmt.Println(i+1,":",temp)
  }
}
