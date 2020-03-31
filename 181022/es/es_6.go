/*
Scrivere un programma Go che legga 4 float
(parametri della retta m e q,
coordinate di un punto P)
e che calcola se il punto P appartiene alla retta
*/

package main

import (
    "fmt"
)

func main() {
  var m,q,x,y float64

  fmt.Print("Inserisci m e q della retta:")
  fmt.Scan(&m,&q)

  fmt.Print("\nInserisci le coordinate del punto:")
  fmt.Scan(&x,&y)

  temp := (x*m + q)
  
  if y == temp{
    fmt.Println("Il punto appartiene alla retta")
  }else{
    fmt.Println("Il punto non appartiene alla retta")
  }
}
