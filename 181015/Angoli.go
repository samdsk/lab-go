/*
Problema Scrivere un programma Go che, date le ampiezze di due angoli di
un triangolo, determini l’ampiezza del terzo angolo.
*/

package main
import "fmt"

func main() {
	var angolo1, angolo2 float64

	fmt.Println("Inserire i due angoli:")
	fmt.Scan(& angolo1, & angolo2)

	fmt.Println("Il terzo angolo ha ampiezza ",
		180-angolo1-angolo2, " gradi.")
}



package main
import "fmt"

func main () {
  var A,B,C,D float64
  LOOP:
  fmt.Println("Questo programma calcola il terzo angolo di un triangolo, conoscendone  due.")
  fmt.Println("Inserisci due angoli, separati da una virgola:")
  fmt.Scanf("%f,%f\n", &A, &B)
  C= 180.0-(A+B)
  D= A+B
  if D>180 {
    fmt.Println("Angoli non validi.")
    goto LOOP
  }
  fmt.Print("Angolo: ")
  fmt.Printf("%.2f", C)
  fmt.Print("\n")
}

package main
import "fmt"

func main() {

	var A1, A2, A3 float64

	fmt.Print("Inserisci il primo angolo: ")
	fmt.Scan(&A1)

	fmt.Print("Inserisci il secondo angolo: ")
	fmt.Scan(&A2)

	A3 = 180-A1-A2

	if A3<0{
	fmt.Print("gli angoli inseriti non sono corretti")
	} else{
	fmt.Print("Il terzo angolo è: ", A3)
	}
	fmt.Scan(&A3)

}
