/*
(A) Calcolare la distanza di Hamming tra due stringhe
*/
package main

import (
    "fmt"
)

func main() {
  var s1,s2 string
  fmt.Print("Inserisci due stringhe: ")
  fmt.Scan(&s1,&s2)

  var hammingDistance int

  for i,k:= range s1 {
    for j,l:= range s2{
      if k!=l && i==j{
        hammingDistance++
      }
    }
  }
  fmt.Println("Le stringhe inserite:")
  fmt.Println(s1)
  fmt.Println(s2)
  fmt.Println("Distanza di Hamming è:",hammingDistance)
}
