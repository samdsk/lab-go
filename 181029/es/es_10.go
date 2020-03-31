/*
10. Scrivere un programma lez3_capitale.go che,
dato un capitale, un tasso di interesse e un obiettivo,
calcola quanti anni occorrono per arrivare all'obiettivo.
*/

package main

import (
    "fmt"
)

func main() {
  var cap,tax,obj,sum float64

  fmt.Print("Inserisci Capitale, Tasso di interesse, Obiettivo: ")
  fmt.Scan(&cap,&tax,&obj)

  i := 1
  for sum < obj{
    sum =+ cap + (cap*tax)/100
    cap = sum
    i++
  }
  
  fmt.Println("Per maturare le interesse fino",obj,"con un tasso di interesse ",tax,"% ci vogliono",i,"anni.")
}
