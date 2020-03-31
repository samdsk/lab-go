/*
Scrivere un programma Go che legga 6 float
(coordinate dei punti A,B,E)
e che calcola se il punto E è equidistante da A e B
*/

package main
import ("fmt"
    "math")

func main () {
  var ax,ay,bx,by,px,py,da,db float64
  
  fmt.Println("Questo programma calcola se un punto P è equidistante da due punti A e B")
  fmt.Println("Fornisci le coordinate del punto A, separate da una virgola: ")
  fmt.Scanf("%f,%f\n", &ax, &ay)
  fmt.Println("Fornisci le coordinate del punto B, separate da una virgola: ")
  fmt.Scanf("%f,%f\n", &bx, &by)
  fmt.Println("Fornisci le coordinate del punto P, separate da una virgola: ")
  fmt.Scanf("%f,%f\n", &px, &py)
  
  da= math.Sqrt((ax-px)*(ax-px)+(ay-py)*(ay-py))
  db= math.Sqrt((bx-px)*(bx-px)+(by-py)*(by-py))
  
  diff=db-da  // e vedere se è più piccolo di un epsilon preso a piacere
  
  if (da==db) {  // funziona sempre???
    fmt.Println("Il punto P è equidistante da A e B.")
  } else {
    fmt.Println("Il punto P non è equidistante da A e B.")
  }
}

/*
package main
import "fmt"
import "math"

func main() {
	var ascissaA, ordinataA, ascissaB, ordinataB, ascissaE, ordinataE, distanzaAE, distanzaBE float64
l
	fmt.Print("Inserire le coordinate dei punti A, B ed E.	"
	fmt.Scan(&ascissaA, &ordinataA, &ascissaB, &ordinataB, &ascissaE, &ordinataE)

	distanzaAE=math.Sqrt((ascissaA-ascissaE)*(ascissaA-ascissaE)+(ordinataA-ordinataE)*(ordinataA-ordinataE))
	distanzaBE=math.Sqrt((ascissaB-ascissaE)*(ascissaB-ascissaE)+(ordinataB-ordinataE)*(ordinataB-ordinataE))

	if distanzaAE==distanzaBE {	//math.Abs(distanzaAE-distanzaBE)<"epsilon"
		fmt.Println("E è equidistante da A e B.")
	} else {
		fmt.Println("E non è equidistante da A e B.")
	}
}
*/

/*
package main
import "fmt"
func main(){

	var ax, bx, ay, by, cx, cy int
	fmt.Println("Digitare cordinate punto A")
	fmt.Scan(&ax,&ay)
	fmt.Println("Digitare cordinate punto B")
	fmt.Scan(&bx,&by)
	fmt.Println("Digitare cordinate punto C da verificare se è equidistante da A e B")
	fmt.Scan(&cx,&cy)
	if (cx-ax)*(cx-ax)+(cy-ay)*(cy-ay)==(cx-bx)*(cx-bx)+(cy-by)*(cy-by){
	fmt.Println("C è equidistante da A e B")
	} else {
	fmt.Println("C non è equidistante da A e B")
	}	
}
*/

/*
package main

import "fmt"
import "math"

func main() {
	var ax, ay, bx, by, ex, ey, distA, distB float64
	fmt.Print("inserisci le coorinate del primo punto ")
	fmt.Scan(&ax, &ay)
	fmt.Print("Inserisci la lunghezza del secondo punto ")
	fmt.Scan(&bx, &by)
	fmt.Print("Inserisci le coordinate del punto da controllare ")
	fmt.Scan(&ex, &ey)
	distA = math.Sqrt(math.Pow(ax-ex, 2) + math.Pow(ay-ey, 2))
	distB = math.Sqrt(math.Pow(bx-ex, 2) + math.Pow(by-ey, 2))
	if math.Abs(distA-distB) <= 0.001 {
		fmt.Println("Il punto è equidistante ")
	} else {
		fmt.Println("Il punto non è equidistante")
	}

}
*/

func main() {

	var Ax, Ay, Ex, Ey, Bx, By float64

	fmt.Print("Inserisci le coordinate di A: ")
	fmt.Scan(&Ax, &Ay)
	fmt.Print("Inserisci le coorsinate di B: ")
	fmt.Scan(&Bx, &By)
	fmt.Print("Inserisci le coordinate di E: ")
	fmt.Scan(&Ex, & Ey)

	if ((Ax-Ex)*(Ax-Ex)+(Ay-Ey)*(Ay-Ey))-((Bx-Ex)*(Bx-Ex)+(By-Ey)*(By-Ey)) < 0.0001 {
		fmt.Print("Equidistante\n")
	}else {
		fmt.Print("Non equidistante\n")
	}

	fmt.Scan(&Ax) //usato come pausa di sistema
}



/*
package main
import "fmt"
import "math"
func main () {
	var ax, ay, bx, by, ex, ey float64
	fmt. Print("punto A: ")
	fmt. Scan(& ax, & ay)
	fmt. Print("punto B: ")
	fmt. Scan(& bx, & by)
	fmt. Print("punto E: ")
	fmt. Scan(& ex, & ey)
	if math. Abs(ex-ax)==math. Abs(ex-bx) && math. Abs(ey-ay)==math. Abs(ey-by){ // mi convince poco, provare con alcuni casi con i punti in diagonale
		fmt. Println("E è equidistante")
	}else{
		fmt. Println("E non è equidistante")
	}
}
*/
