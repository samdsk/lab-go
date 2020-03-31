package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Fraction struct {
	num, den int
}

func main() {
	if len(os.Args) != 5 {
		fmt.Println("numero di parametri non valido")
		os.Exit(0)
	}
	n1, _ := strconv.Atoi(os.Args[1])
	d1, _ := strconv.Atoi(os.Args[2])
	n2, _ := strconv.Atoi(os.Args[3])
	d2, _ := strconv.Atoi(os.Args[4])

	f1, err1 := newFraction(n1, d1)
	f2, err2 := newFraction(n2, d2)
	if err1 != nil || err2 != nil {
		fmt.Println("divisore nullo")
	} else {
		//fmt.Println(f1, f2)
		f := sum(f1, f2)
		fmt.Println(f.num, "/", f.den)
	}
}

func newFraction(n, d int) (f Fraction, err error) {
	if d == 0 {
		err = errors.New("divisore nullo")
		return
	}
	//riduzione ai minimi termini
	min := int(math.Min(float64(n), float64(d)))
	for i := min; i >= 2; i-- {
		if n%i == 0 && d%i == 0 {
			n /= i
			d /= i
		}
	}
	f.num = n
	f.den = d
	return
}

func sum(f1, f2 Fraction) (sum Fraction) {
	d := f1.den * f2.den
	n := f1.num*f2.den + f2.num*f1.den
	sum, _ = newFraction(n, d)
	return
}
