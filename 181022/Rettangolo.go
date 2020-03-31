/*
Scrivere un programma Go che legga sei interi
(coordinate dei punti A e B diagonali di un
rettangolo e coordinate di un punto P)
e che calcola se il punto P è fuori
o dentro il rettangolo AB


Esempio:

A (1,1)
B (2,2)

P (5,5)

è fuori

etc.

*/

/*
package main
import "fmt"
func main() {
	var ax, ay, bx, by, px, py int
	
	fmt.Println("inserisci prima i punti x e y del punto in basso a sinistra, i punti x e y del punto in alto a destra e i punti del punto scelto")
	
	fmt.Scan(&ax, &ay, &bx, &by, &px, &py)
		if ax < bx && ay < by {
	
	if (px <= bx && px >= ax) && (py <= by && py >= ay) {
		fmt.Println("il punto è contenuto")
	} else {
		fmt.Println("il punto non è contenuto")
	}
	}
}
*/

package main
import "fmt"

func main() {
	var ascissaA, ordinataA, ascissaB, ordinataB, ascissaP, ordinataP int

	fmt.Print("Inserire le coordinate del punto A.	")
	fmt.Scan(&ascissaA, &ordinataA)
	fmt.Print("Inserire le coordinate del punto B.	")
	fmt.Scan(&ascissaB, &ordinataB)
	fmt.Print("Inserire le coordinate del punto P.	")
	fmt.Scan(&ascissaP, &ordinataP)

	if (ascissaP<ascissaA || ascissaP>ascissaB || ordinataP<ordinataA || ordinataP>ordinataB)  {
		fmt.Println("Il punto non appartiene al rettangolo.")
	} else {
		fmt.Println("Il punto appartiene al rettangolo.")
	}
}


/*

package main
import "fmt"
func main() {
  var xA, yA, xB, yB, xP, yP int
  fmt.Println("inserire coordinate (x,y) di A\n")
  fmt.Scan(&xA, &yA)
  fmt.Println("inserire coordinate (x,y) di B\n")
  fmt.Scan(&xB, &yB)
  fmt.Println("inserire coordinate (x,y) di P\n")
  fmt.Scan(&xP, &yP)

  if ((xP>=xA && xP<=xB)||(xP>=xB && xP<=xA)) && ((yP>=yA && yP<=yB)||(yP>=yB && yP<=yA)){
  fmt.Printf("il punto di coordinate (%d,%d) e' all'interno del quadrato\n",xP,yP)
  }else{
  fmt.Printf("il punto di coordinate (%d,%d) non e' all'interno del quadrato\n",xP,yP)
  }
}



*/
