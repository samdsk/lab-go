/*
Problema Scrivere un programma Go che, date le misure dell’altezza e della
base di un rettangolo, calcoli il perimetro e l’area.
*/

package main

import (
    "fmt"
)

func main() {
  var base, altezza float64
  fmt.Println("Inserisci base e altezza ")
  fmt.Scan(&base, &altezza)
  fmt.Println("Il perimetro è: ", 2*(base+altezza), " L'area: ", base*altezza)  
}
