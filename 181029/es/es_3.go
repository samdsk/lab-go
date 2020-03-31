/*
3. Scrivere un programma lez3_somma.go che,
dato un intero n, somma i numeri da n a 1,
stampando le somme parziali.
*/


package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Print("Inserisci n: ")
    fmt.Scan(&n)

    sum := 0

    for i := 1; i <= n; i++ {
      sum += i
      fmt.Println(i,"\t- Sum:\t", sum)
    }
}
