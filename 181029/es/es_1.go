/*
1. Scrivere un programma lez3_asterischi.go che, dato un intero n, stampa n asterischi.
*/

package main

import (
    "fmt"
)

func main() {
    var i int;
    fmt.Print("Inserisci n: ")
    fmt.Scan(&i)

    for x := 0; x<i; x++{
      fmt.Print("*")
    }

    fmt.Print("\n")
}
