/*
Problema Scrivere un programma Go che calcola la distanza tra due punti
nel piano cartesiano.

N.B. usare "math" (cfr. https://golang.org/pkg/)
*/

package main
import ("fmt"
        "math")

func main () {
  var x,y,q,z,D float64
  fmt.Println("Questo programma calcola la distanza fra due punti del piano cartesiano.")
  fmt.Println("Inserisci le coordinate del primo punto, separate da una virgola:")
  fmt.Scanf("%f,%f\n", &x, &y)
  fmt.Println("Inserisci le coordinate del secondo punto, separate da una virgola:")
  fmt.Scanf("%f,%f\n", &q, &z)
  D= math.Sqrt(((x-q)*(x-q))+((y-z)*(y-z)))
  fmt.Println("Distanza:", D)
}



package main

import "fmt"
import "math"

func main() { 	
	var xa, xb, ya, yb, xip, yip, ipotenusa float64
	fmt.Print("inserisci le coordinate del primo punto (x y)")
	fmt.Scan(&xa, &ya)
	fmt.Print("inserisci le coordinate del secondo punto (x y)")
	fmt.Scan(&xb, &yb)
	xip = math.Abs(xa - xb)
	yip = math.Abs(ya - yb)
	ipotenusa = math.Sqrt((yip * yip) + (xip * xip))
	fmt.Print("distanza=", ipotenusa)
}


package main

import (
	"fmt"
	"math"
)

func main() {
	var ascissa1, ascissa2, ordinata1, ordinata2, distanza float64
	var quadrato1, quadrato2 float64
	fmt.Println("Inserisci coordinate primo punto")
	fmt.Scan(&ascissa1, &ordinata1)
	fmt.Println("Inserisci coordinate secondo punto")
	fmt.Scan(&ascissa2, &ordinata2)
	quadrato1 = math.Pow(ascissa1-ascissa2, 2)
	quadrato2 = math.Pow(ordinata1-ordinata2, 2)
	distanza = math.Sqrt(quadrato1 + quadrato2)
	fmt.Println("La distanza tra i due punti è:", distanza)
}
