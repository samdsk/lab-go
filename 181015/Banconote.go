	
/*
Problema: Scrivere un programma Go che, dato un ammontare (valore intero)
di denaro, lo suddivide nel numero più piccolo possibile di banconote.

Annotazioni: Si supponga che i tagli di banconote disponibili siano: 100, 50,
20, 10, 5, 2, 1.
*/

//Se nella shell metto echo 1000 | ./Banconote mi assegna il valore

package main

import "fmt"

func main () {
  var n int
  fmt.Println("Questo programma divide un ammontare di denaro nel numero più piccolo possibile di banconote.")
  fmt.Println("Inserisci l'ammontare:")
  fmt.Scan(&n)
  fmt.Println("Banconote da 100:", n/100)
  n=n%100
  fmt.Println("Banconote da 50:", n/50)
  n=n%50
  fmt.Println("Banconote da 20:", n/20)
  n=n%20
  fmt.Println("Banconote da 10:", n/10)
  n=n%10
  fmt.Println("Banconote da 5:", n/5)
  n=n%5
  fmt.Println("Banconote da 2:", n/2)
  n=n%2
  fmt.Println("Banconote da 1:", n)
}
