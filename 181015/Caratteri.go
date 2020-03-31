package main

import "fmt"

func main() {
	var carattere rune
	carattere=65  // ASCII 'A'
	fmt.Printf("carattere é di tipo: %T\n", carattere)
	//Stampa int32^^^
	fmt.Println(string(carattere))
	carattere++  // diventa ASCII 'B'
	fmt.Println(string(carattere))
}
