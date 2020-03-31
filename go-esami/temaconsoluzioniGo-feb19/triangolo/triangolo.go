package main

import (
	"fmt"
)

type Punto struct {
	x, y float64
}

type Retta struct {
	m, q float64
}

func main() {
	var m1, q1, m2, q2 float64
	var x, y float64

	fmt.Scan(&m1, &q1)
	r1 := Retta{m1, q1}

	fmt.Scan(&m2, &q2)
	r2 := Retta{m2, q2}

	fmt.Scan(&x, &y)
	p := Punto{x, y}

	if p.y >= 0 && Sotto(p, r1) && Sotto(p, r2) {
		fmt.Println(p, "dentro")
	} else {
		fmt.Println(p, "fuori")
	}
}

func Sotto(p Punto, r Retta) bool {
	return p.y <= p.x*r.m+r.q
}
