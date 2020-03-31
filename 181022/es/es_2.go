/*
Scrivere un programma Go che prenda in ingresso un float64
e determini se esso sia anche un int.
(Esempio: 5.0 è anche un intero, mentre 5.31 no)

(Opzionale: se la risposta è no arrotondarlo
all'intero più vicino, senza usare il package math)
*/

package main

import (
    "fmt"
)

func main() {
    var f float64
    var i int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&f)

    i = int(f)
	var temp float64
	temp = f - float64(i)

    if temp != 0{
      fmt.Println("Il numero inserito è un float")

      if temp >= 0.5 {
        i++
        fmt.Println("Il numero è stato arrotondato per eccesso: ", i)
      }else{
		fmt.Println("Il numero è stato arrotondato per difetto: ", i)
	  }
    }else{
      fmt.Println("Il numero inserito è un intero")
      fmt.Println("Il numero è stato arrotondato per difetto: ", i)
    }

}
