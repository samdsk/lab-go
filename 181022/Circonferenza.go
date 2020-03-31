/*
Scrivere un programma Go che legga 5 interi (e versione float)
(coordinate dei punti C del centro di un
cerchio, raggio e coordinate di un punto P)
e che calcola se il punto P è fuori
o dentro il cerchio
*/

/* non corretto per float, considera quadrato che circoscrive
package main
import "fmt"
func main() {
	var xc, yc, r, xp, yp float64
	fmt.Print("inserisci le coordinate del cerchio ")
	fmt.Scan(&xc, &yc)
	fmt.Print("inserisci il raggio ")
	fmt.Scan(&r)
	fmt.Print("inserisci le coordinate del punto ")
	fmt.Scan(&xp, &yp)
	if (xc-r <= xp) && (yc-r <= yp) && (xc+r >= xp) && (yc+r >= yp) {
		fmt.Print("Il punto è contenuto")
	} else {
		fmt.Print("Il punto non è contenuto")
	}
}
*/

package main
import "fmt"
import "math"

func main() { 
	var ascissaC, ordinataC, raggio, ascissaP, ordinataP, distanza float64
	

	fmt.Println("Inserire le coordinate del centro del cerchio, il suo raggio, e le coordinate di P.")
	fmt.Scan(&ascissaC, &ordinataC, &raggio, &ascissaP, &ordinataP)

	distanza=math.Sqrt((ascissaC-ascissaP)*(ascissaC-ascissaP)+(ordinataC-ordinataP)*(ordinataC-ordinataP))

	if distanza<=raggio {
		fmt.Println("Il punto appartiene al cerchio.")
	} else {
		fmt.Println("Il punto non appartiene al cerchio.")
	}
}


/*
package main
import "fmt"
import "math"

func main() {
	var cx, cy, r, px, py, dist float64
	fmt.Print("inserisci le coorinate del centro ")
/	fmt.Scan(&cx, &cy)
	fmt.Print("Inserisci la lunghezza del raggio ")
	fmt.Scan(&r)
	fmt.Print("Inserisci le coordinate del punto da controllare ")
	fmt.Scan(&px, &py)
	dist = math.Sqrt(math.Pow(cx-px, 2) + math.Pow(cy-py, 2))
	if dist < r {
		fmt.Print("Il punto è dentro il cerchio ")
	} else {
		fmt.Println("Il punto è fuori dal centro")
	}

}
*/




/*
package main
import "fmt"
func main () {
	var cx, cy, r, px, py int
	fmt. Print("Coordinate centro cerchio: ")
	fmt. Scan(& cx, & cy)
	fmt. Print("Lunghezza raggio: ")
	fmt. Scan(& r)
	fmt. Print("Coordinate P: ")
	fmt. Scan(& px, & py)
	if px>cx+r && py>cy+r || px<cx-r && px<cx-r{
		fmt. Println("P è fuori")
	}else{
		fmt. Println("P è dentro")
	}
}



*/

/*
package main
import "fmt"
func main(){

	var cx,cy,r,px,py int
	fmt.Println("Inserire la lunghezza del raggio")
	fmt.Scan(&r)
	fmt.Println("Inserire le coordinate del centro")
	fmt.Scan(&cx,&cy)
	fmt.Println("Inserire le coordinate del punto")
	fmt.Scan(&px,&py)
	if (px-cx)*(px-cx)+(py-cy)*(py-cy)>r*r {
		fmt.Println("il punto non appartiene al cerchio")
	} else {
	fmt.Println("il punto appartiene al cerchio")
	}		
}
*/








