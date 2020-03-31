/*
2. Scrivere un programma lez3_euclide.go che, dati due interi, calcolare il MCD tra i due numeri con l'algoritmo di Euclide.
*/

package main

import (
    "fmt"
    "math"
)

func main() {
    var a,b,r int
    fmt.Print("Inserisci due numeri a e b per calcolare MCD: ")
    fmt.Scan(&a,&b)

    a = int(math.Abs(float64(a)))
    b = int(math.Abs(float64(b)))

    if a < b {
      temp := a
      a = b
      b = temp
    }

    if b == 0{
      fmt.Println("MCD è: ",a)
      return
    }

    for {
      if b == 0 {
        break
      }
      r = a%b

      a = b
      b = r
    }

    fmt.Println("MCD è: ",a)

}
