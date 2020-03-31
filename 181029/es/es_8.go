/*
8. Scrivere un programma lez3_andamento.go che legge da tastiera una serie di numeri > -1
e stampa "andamento positivo" ogni volta che il nuovo
valore e` maggiore o uguale al precedente e "andamento negativo" altrimenti.
Si ferma quando il numero in input e` -1.
*/

package main

import (
    "fmt"
)

func main() {

  temp := .0

    for{
      var a float64

      fmt.Print("Inserisci un numero > -1 (-1 per uscire dal prog.): ")
      fmt.Scan(&a)

      if a >= temp{
        fmt.Println("\nAndamento positivo\n\n")
      }else if a == -1{
        break
      }else{
        fmt.Println("\nAndamento negativo\n\n")
      }

      temp = a
    }
}
