/*
Scrivere un programma Go che legga 5 float
(parametri della retta m e q,
coordinate di un punto P
e un valore float per la distanza D)
e che calcola se il punto P è a meno di una distanza D dalla retta
*/

package main

import (
    "fmt"
    "math"
)

func main() {
  var m,q,x,y,d float64

  fmt.Print("Inserisci m e q della retta:")
  fmt.Scan(&m,&q)

  fmt.Print("\nInserisci le coordinate del punto:")
  fmt.Scan(&x,&y)

  fmt.Print("\nInserisci la distanza:")
  fmt.Scan(&d)

  //retta perpendicolare alla retta data passante per il punto P (formula esplicità)

  temp := math.Abs(y - (m*x + q))/ math.Sqrt(1+m*m)

  if temp <= d{
    fmt.Println("Il punto è interno alla distanza ", d)
  }else{
    fmt.Println("Il punto non è interno alla distanza ", d)
  }

}
