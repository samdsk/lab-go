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

		/*
			if err != nil {
				return
			}
		*/

		//controllo parità
		if (previous < 0 && follower < 0) || (previous >= 0 && follower >= 0) || err != nil {
			fmt.Println("elemento in posizione", i, "non valido")
			return
		}
		previous = follower
	}
	fmt.Println("sequenza valida")
}
