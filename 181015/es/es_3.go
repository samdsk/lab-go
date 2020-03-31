/*
Problema: Scrivere un programma Go che, dato un ammontare (valore intero)
di denaro, lo suddivide nel numero più piccolo possibile di banconote.

Annotazioni: Si supponga che i tagli di banconote disponibili siano: 100, 50,
20, 10, 5, 2, 1.
*/

package main

import (
    "fmt"
)

func main() {
  var a int32
  fmt.Println("Inserisci ammontare del denaro:")
  fmt.Scan(&a)
  fmt.Println("Inserito: ", a)
  fmt.Println("Banconote da 100:", a/100)
  a = a%100
  fmt.Println("Banconote da 50:", a/50)
  a = a%50
  fmt.Println("Banconote da 20:", a/20)
  a = a%20
  fmt.Println("Banconote da 10:", a/10)
  a = a%10
  fmt.Println("Banconote da 5:", a/5)
  a = a%5
  fmt.Println("Monete da 2:", a/2)
  a = a%2
  fmt.Println("Monete da 1:", a/1)  
}
