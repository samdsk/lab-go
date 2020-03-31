/*
12. Scrivere un programma lez3_pos_cifra.go che legge da tastiera
un numero intero n di quattro cifre e un numero i di una cifra
e stampa in che posizione (1 per le unita`) e` la prima occorrenza di
i nella sequenza di cifre che rappresentano n (0 se i non e` presente).
*/

package main

import (
  "fmt"
  "math"
)

func main() {
  var n,i int
  fmt.Print("Inserisci un numero con 4 cifre: ")
  fmt.Scan(&n)
  fmt.Print("\nInserisci un numero da controllare:")
  fmt.Scan(&i)



  for k:=1;k<=4;k++{
    temp := n/int(math.Pow(10.,float64(4-k)))
    n = n%int(math.Pow(10.,float64(4-k)))
    if i == temp {
      fmt.Print(1)
    }else{
      fmt.Print(0)
    }
  }
}
