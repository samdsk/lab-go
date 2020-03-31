/*
Scrivere un programma "Clessidra" che prenda in ingresso un intero n e faccia il conto alla rovescia da n a 0 in secondi.
Hint: usare le funzioni Now, Sub e Seconds del package time;
non usare Sleep!
*/

package main

import (
  "fmt"
  "time"
)

func main() {
  var n int
  fmt.Print("Inserisci numero di secondi: ")
  fmt.Scan(&n)

  for n>0{
    start := time.Now()

    for {
      t := time.Now().Sub(start).Seconds();
      if t >= 1{break}
    }

    // Versione publicata sul sito
    // for time.Now().Sub(start).Seconds() < 1 {}

    fmt.Println("Seconds:",n)
    n--
  }
}
