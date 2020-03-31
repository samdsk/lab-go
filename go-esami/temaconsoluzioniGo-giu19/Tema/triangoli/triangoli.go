package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

type Punto struct {
	x, y float64
}

type Triangolo struct {
	A, B, C Punto
}

func main() {
	var P1, P2, P3 Punto
	fmt.Scan(&P1.x, &P1.y)
	fmt.Scan(&P2.x, &P2.y)
	fmt.Scan(&P3.x, &P3.y)

	triangolo, err := newTriangolo(P1, P2, P3)

	fmt.Println(lunghezzeLati(P1, P2, P3))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("triangolo", triangolo)
		fmt.Println("tipo", tipoTriangolo(triangolo))
	}
}

func distanza(P1, P2 Punto) float64 {
	dist := math.Sqrt(math.Pow(P1.x-P2.x, 2) + math.Pow(P1.y-P2.y, 2))
	return dist
}

func lunghezzeLati(A, B, C Punto) (dist [3]float64) {
	dist[0] = distanza(A, B)
	dist[1] = distanza(B, C)
	dist[2] = distanza(C, A)
	return
}

func newTriangolo(A, B, C Punto) (triangolo Triangolo, err error) {
	dist := lunghezzeLati(A, B, C)
	if dist[0]+dist[1] > dist[2] && dist[1]+dist[2] > dist[0] && dist[0]+dist[2] > dist[1] {
		triangolo = Triangolo{A, B, C}
	} else {
		err = errors.New("non è un triangolo")
	}
	return
}

func tipoTriangolo(triangolo Triangolo) int {
	epsilon := 1e-6
	dist := lunghezzeLati(triangolo.A, triangolo.B, triangolo.C)
	sort.Float64s(dist[:])
	if math.Abs(dist[0]-dist[1]) < epsilon && math.Abs(dist[1]-dist[2]) < epsilon {
		return 3
	} else if math.Abs(dist[0]-dist[1]) > epsilon && math.Abs(dist[1]-dist[2]) > epsilon {
		return 0
	} else {
		return 2
	}
}
