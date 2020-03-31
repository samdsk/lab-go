/*
B) Calcolare la distanza di ...(**) tra due stringhe di = lunghezza

(**) Somma della distanza (valori assoluti) dei caratteri nella stessa posizione

ciao
ciap
0001

a
A
97-65
*/

package main

import (
    "fmt"
    "math"
)

func main() {
  var s1,s2 string
  fmt.Print("Inserisci due stringhe: ")
  fmt.Scan(&s1,&s2)

  sum := 0

  for i,k:= range s1 {
    for j,l:= range s2{
      if i == j{
        sum += int(math.Abs(float64((int(k)-int(l)))))
      }
    }
  }

  fmt.Println("La somma della distanza:",sum)

}
