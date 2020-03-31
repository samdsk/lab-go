/*
Problema Scrivere un programma Go che, date le misure dell’altezza e della
base di un rettangolo, calcoli il perimetro e l’area.
*/

package main
import "fmt"

func main() {
	var base, altezza float64

	fmt.Println("Inserire base e altezza:")
	fmt.Scan(&base, &altezza)

	fmt.Println("Il perimetro è", (base+altezza)*2)
	fmt.Println("L'area è", base*altezza)
}

package main
import "fmt"

func main() {

	var B, H, A, _2p float64

	fmt.Print("Inserisci il'altezza del rettangolo: ")
	fmt.Scan(&H)

	fmt.Print("Inserisci la base del rettangolo: ")
	fmt.Scan(&B)

	A = B*H
	_2p = 2*B + 2*H

		fmt.Print("\nl'area è: ",A)
		fmt.Print("\nIl perimetro è: ", _2p)

	fmt.Scan(&_2p)

}


package main

import "fmt"

func main() {
	var base, altezza, area, perimetro float64
	fmt.Println("Inserisci base ed altezza:")
	fmt.Scan(&base, &altezza)
	perimetro = base*2 + altezza*2
	area = base * altezza
	fmt.Println("L'area del rettangolo è", area, " e il perimetro", perimetro)
}
