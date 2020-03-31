package main

import (
	"esami"	
	"fmt"
)

func main() {
	p := make(esami.Nomi2Voti)

	p["Paolo Boldi"] = []int{18, 24, 25, 19}
	p["Giovanni Rossi"] = []int{30, 32}
	p["Marco Bianchi"] = []int{18, 32}

	m := esami.Medie(p)
	for nome, media := range m {
		fmt.Printf("%-30s\t%.2f\t%.2f\n", nome, media, esami.Trentesimi2Centodecimi(media))
	}

	esami.ScriviFile(p, "pippo.txt")
}
	