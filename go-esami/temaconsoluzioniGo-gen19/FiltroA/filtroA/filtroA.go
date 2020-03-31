package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	len := len(os.Args)

	if len == 1 {
		fmt.Println("nessun valore in ingresso")
		return
	}

	previous, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("elemento in posizione 1 non valido")
		return
	}

	for i := 2; i < len; i++ {
		follower, err := strconv.Atoi(os.Args[i])

		//controllo parità: diverse alternative
		//1.
		/* if (follower%2 == 0 && previous%2 == 0) || (follower%2 != 0 && previous%2 != 0) {
		//al posto dell'if qui sotto
		*/

		//2.
		if (follower+previous)%2 == 0 || err != nil {
			fmt.Println("elemento in posizione", i, "non valido")
			return
		}
		previous = follower

		//3.
		/*
			bit := (previous) % 2 //istruzione da mettere prima del for (riga 22)

			if follower%2 == bit || err != nil { //in questo caso non servono due var distinte prev e follow
				//ma serve var bit in più
				fmt.Println("elemento in posizione", i, "non valido")
				return
			}
			bit = (bit + 1) % 2
		*/

	}
	fmt.Println("sequenza valida")
}
