/*
Scrivere un programma Go che legga 4 float
(parametri della retta m e q,
coordinate di un punto P)
e che calcola se il punto P appartiene alla retta
*/

package main
import "fmt"
func main () {
  var m,q,px,py float64
  fmt.Println("Questo programma calcola se un punto P appartiene a una retta nella forma y=mx+q")
  fmt.Println("Fornisci le coordinate del punto P, separate da una virgola: ")
  fmt.Scanf("%f,%f\n", &px, &py)
  fmt.Println("Fornisci la pendenza m della retta: ")
  fmt.Scan(&m)
  fmt.Println("Fornisci il parametro q: ")
  fmt.Scan(&q)
  if py==m*px+q { // qui problema?
    fmt.Println("Il punto P appartiene alla retta.")
  } else {
    fmt.Println("Il punto P non appartiene alla retta.")
  }
}

/*
package main
import "fmt"
func main() {
	var m, q, ascissaP, ordinataP float64

	fmt.Println("Inserire i parametri m e q della retta e le coordinate del punto P.")
	fmt.Scan(&m, &q, &ascissaP, &ordinataP)

	if m*ascissaP+q==ordinataP {
		fmt.Println("Il punto appartiene alla retta.")
	} else {
		fmt.Println("Il punto non appartiene alla retta.")
	}
}
*/
