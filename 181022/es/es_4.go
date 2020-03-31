/*
Scrivere un programma Go che legga 5 interi
(coordinate dei punti C del centro di un
cerchio, raggio e coordinate di un punto P)
e che calcola se il punto P è fuori
o dentro il cerchio
*/

package main
import (
    "fmt"
    "math"
)

func main() {
  var x,y,r,xp,yp int

  fmt.Print("Inserisci coordinate del centro: ")
  fmt.Scan(&x,&y)

  fmt.Print("\nInserisci raggio: ")
  fmt.Scan(&r)

  fmt.Print("\nInserisci coordinate del puunto P: ")
  fmt.Scan(&xp, &yp)

  d := math.Sqrt(math.Pow(float64(x-xp),2) + math.Pow(float64(y-yp),2))

  if d <= float64(r) {
    fmt.Println("Il punto P è interno alla circonferenza")
  }else{
    fmt.Println("Il punto P è esterno alla circonferenza")
  }

}
