/*
Problema Scrivere un programma Go che, date le ampiezze di due angoli di
un triangolo, determini l’ampiezza del terzo angolo.
*/

package main

import (
    "fmt"
)

func main() {
  var angoloA, angoloB float64
  fmt.Println("Inserisci i due angoli: ")
  fmt.Scan(&angoloA, &angoloB)
  fmt.Println("Il terzo angolo è: ", 180 - angoloA - angoloB)
}
