/*
Scrivere un programma Go che legga 5 float
(parametri della retta m e q,
coordinate di un punto P
e un valore float per la distanza D)
e che calcola se il punto P è a meno di una distanza D dalla retta
*/

package main
import ("fmt"
        "math")
func main () {
  var m,q,d,px,py float64
  fmt.Println("Questo programma calcola se un punto P si trova a meno di una distanza D da una retta nella forma y=mx+q")
  fmt.Println("Fornisci le coordinate del punto P, separate da una virgola: ")
  fmt.Scanf("%f,%f\n", &px, &py)
  fmt.Println("Fornisci la distanza d dalla retta: ")
  fmt.Scan(&d)
  fmt.Println("Fornisci la pendenza m della retta: ")
  fmt.Scan(&m)
  fmt.Println("Fornisci il parametro q: ")
  fmt.Scan(&q)
  // provato con 5,5 dist=3, retta y=4, NON funziona
  if math.Abs(py-m*px+q)<=d {
    fmt.Println("Il punto P rientra nella distanza D.")
  } else {
    fmt.Println("Il punto P non rientra nella distanza D.")
  }
}

/*
package main
import (
	"fmt"
 	"math"
)

func main() {
	var m, q, ascissaP, ordinataP, distanzaP float64

	fmt.Println("Inserire i parametri m e q della retta, le coordinate del punto P e la distanza presunta.")
	fmt.Scan(&m, &q, &ascissaP, &ordinataP, &distanzaP)

	if distanzaP>math.Abs(ordinataP-(m*ascissaP+q))/math.Sqrt(1+math.Pow(m, 2)) {
		fmt.Println("La distanza inserita è maggiore della reale.")
	} else if distanzaP==math.Abs(ordinataP-(m*ascissaP+q))/math.Sqrt(1+math.Pow(m, 2)) {
		fmt.Println("La distanza inserita è uguale alla reale.")
	} else  {
		fmt.Println("La distanza inserita è minore della reale.")
	}
}
*/

/*
package main

import "fmt"
import "math"

func main() {
	var ax, ay, bx, by, m1, q1, m2, q2, d, dist float64
	
	fmt.Print("Inserisci le coordinate del punto ")
	fmt.Scan(&ax,&ay)
	fmt.Print("Inserisci l'm della retta poi il q della retta ")
	fmt.Scan(&m1,&q1)
	fmt.Print("Inserisci la distanza ")
	fmt.Scan(&d)

	m2 = -1 / m1
	q2 = -m2*ax + ay
	bx = (q2 - q1) / (m1 - m2)
	by = bx*m2 + q2
	dist = math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	if dist < d {
		fmt.Println("il punto è a distanza minore di ", d)
	} else {
		fmt.Println("il punto è a distanza maggiore di ", d)
	}
}
*/


/*
package main
import "fmt"
import "math"

func main () {
	var  m, q, xp, yp, D float64
	fmt.Println("Inserisci le coordinate del punto P: ")
	fmt.Scan(&xp)
	fmt.Scan(&yp)
	fmt.Println("Inserisci la distanza dalla retta: ")
	fmt.Scan(&D)
	fmt.Println("Inserisci la pendenza m della retta(y=mx+q): ")
	fmt.Scan(&m)
	fmt.Println("Inserisci il parametro q della retta(y=mx+q): ")
	fmt.Scan(&q)
	if (math.Abs(yp-(m*xp+q)))/math.Sqrt(1+m*m)<= D{
		fmt.Println("Il punto P è dentro la distanza fornita.")
	} else {
		fmt.Println("Il punto P non è dentro la distanza fornita.")
	}
}
*/
