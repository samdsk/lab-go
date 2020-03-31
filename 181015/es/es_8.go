/*
	Problema:
	scrivere un programma che calcoli (approssimativamente,
	non importano i bisestili) la data corrispondente
	ad un numero di secondi trascorsi da "the epoch" (1 gen 1970)

	es. @2147483647  che giorno è?
*/

package main

import (
    "fmt"
)

func main() {
  var a, day int64
	fmt.Println("Inserisci la data in secondi (epoch):")

	fmt.Scan(&a)
	fmt.Println("Inserito: ", a)

	day = 3600*24

	fmt.Println("Anni: ", a/(day*365), " Anno: ", 1970 + a/(day*365))
	a = a%(day*365)
	fmt.Println("Mesi: ", a/(day*30))
	a = a%(day*30)
	fmt.Println("Giorni: ", a/day)
}
