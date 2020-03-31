package main

import (
	"fmt"
	"sort"
)

func main() {
	var x float64
	var list []float64
	for {
		_, err := fmt.Scan(&x)
		if err != nil {
			break
		}
		list = append(list, x)
	}
	moda := moda(list)
	fmt.Println("moda:", moda)
	fmt.Println("mediana:", mediana(list))
}

func moda(serie []float64) (mode float64) {
	max := 0
	occorrenze := make(map[float64]int)
	for _, x := range serie {
		occorrenze[x]++
		if max < occorrenze[x] {
			max = occorrenze[x]
		}
	}

	//Se c'è più di un valore con la stessa frequenza max, restituisce il più piccolo
	keys := make([]float64, 0, len(occorrenze))
	for k, _ := range occorrenze {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	for _, k := range keys {
		if occorrenze[k] == max {
			mode = k
			break
		}
	}
	return
}

func mediana(serie []float64) (mediana float64) {
	l := len(serie)
	sort.Float64s(serie)
	if l%2 == 1 {
		mediana = float64(serie[l/2])
	} else {
		mediana = float64(serie[l/2]+serie[l/2-1]) / 2.
	}
	return
}
