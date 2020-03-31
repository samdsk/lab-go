package main

import (
	"fmt"
	"io"
	"math"
)

type Punto struct {
	x float64
	y float64
}

func main() {
	var x, y float64
	var tappe []Punto

	for {
		_, err := fmt.Scan(&x, &y)
		if err == io.EOF {
			break
		}
		p := Punto{x, y}
		tappe = append(tappe, p)
	}
	fmt.Println("tappe:", tappe)
	fmt.Println("percorso", percorso(tappe))
}

func distanza(p1, p2 Punto) (d float64) {
	x := p1.x - p2.x
	y := p1.y - p2.y
	d = math.Sqrt(x*x + y*y)
	return
}

func percorso(tappe []Punto) (tot float64) {
	tot = 0
	lung := len(tappe)
	for i := 1; i < len(tappe); i++ {
		tot += distanza(tappe[i], tappe[i-1])
	}
	tot += distanza(tappe[lung-1], tappe[0])
	return
}
