/*
7. Scrivere un programma lez3_pari.go che conta il numero di
numeri pari in 20 numeri interi generati in modo casuale.
Usare la funzione rand.Int() (o rand.Intn(n int)) del package "math/rand";
per generare sequenze sempre diverse, importare il package "time" e usare l'istruzione

rand.Seed(time.Now().UnixNano())

prima di iniziare a generare numeri (ad esempio come prima istruzione del main).
*/
/*
7.b. Modificare il programma precedente in modo che i numeri siano compresi tra 1 e 10.
*/
package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
  var n,i int
  rand.Seed(time.Now().UnixNano())
  
  for i < 20{
    n = rand.Intn(20)
    if temp := n%2; temp == 0{
      fmt.Println(i," - ",n)
      i++
    }
  }

}
