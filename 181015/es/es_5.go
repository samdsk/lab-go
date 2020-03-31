/*
Problema Scrivere un programma Go che calcola la distanza tra due punti
nel piano cartesiano.

N.B. usare "math" (cfr. https://golang.org/pkg/)
*/

package main

import (
    "fmt"
)

func main() {
  var a, b, c, d, pX, pY float64

  fmt.Println("Inserisci le coordinate del primo punto:")
  fmt.Scan(&a, &b)

  fmt.Println("Inserisci le coordinate del secondo punto:")
  fmt.Scan(&c, &d)

  pX = a-c
  pY = b-d

  fmt.Println("La distanza è: ", math.Sqrt((pX*pX)+(pY*pY)))  
}
