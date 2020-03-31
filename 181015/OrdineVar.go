package main

import "fmt"

	var a = b+c		//Variabili globali
	//var b = 7
	var b = a
	var c = 1

func main() {

	/* poi provare anche qui dentro
	var a = b+c
	//var b = 7
	var b = a
	var c = 1
	*/


	fmt.Println("ciao")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
