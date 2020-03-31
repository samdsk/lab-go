/*
6. Scrivere un programma lez3_indovina.go
che fissa un numero intero e chiede di provare a
indovinare il numero fissato fino a indovinarlo.
*/

package main

import (
    "fmt"
    "math/rand"
)

func main() {
    var n int

    n = rand.Intn(30)
    fmt.Println("Prova indovinare il numero nascosto (Hint: 0<=n<=30)")
    for{
      temp := 0
      fmt.Println("\nInserisci un numero: ")
      fmt.Scan(&temp)

      if temp == n{
        fmt.Println("Hai indovinato il numero nascosto!!!")
        break
      }else if temp > n{
        fmt.Println("(Hint: Il numero è minore di quello che hai inserito)")
      }else{
        fmt.Println("(Hint: Il numero è maggiore di quello che hai inserito)")
      }
    }
}
