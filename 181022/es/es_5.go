/*
Scrivere un programma Go che legga 6 float
(coordinate dei punti A,B,E)
e che calcola se il punto E è equidistante da A e B
*/

package main

import (
    "fmt"
    "math"
)

func main() {
  var xA,xB,xE,yA,yB,yE float64

  fmt.Print("Inserisci coordinate del punto A: ")
  fmt.Scan(&xA,&yA)

  fmt.Print("\nInserisci coordinate del punto B: ")
  fmt.Scan(&xB,&yB)

  fmt.Print("\nInserisci coordinate del punto E: ")
  fmt.Scan(&xE,&yE)

  dA := math.Sqrt(math.Pow(xA-xE, 2) + math.Pow(yA-yE, 2))
  dB := math.Sqrt(math.Pow(xB-xE, 2) + math.Pow(yB-yE, 2))

  if dA == dB {
    fmt.Println("Il punto E è equidistante da A e B")
  }else{
    fmt.Println("Il punto E non è equidistante da A e B")
  }

}
