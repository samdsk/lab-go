/*
Problema Scrivere un programma Go che converta in Fahrenheit
una temperatura espressa in gradi centigradi.
*/

package main

import (
  "fmt"
)

func main() {
  var celsius, farenheit float64
  fmt.Println("Inserisci la temperatura in °C da convertire in °F")
  fmt.Scan(&celsius)

  farenheit  = (celsius * 9/5) + 32

  fmt.Println("La temperatura inserita °C: ", celsius, " in Farenheit °F: ", farenheit )
}
