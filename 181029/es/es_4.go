/*
4. Scrivere un programma lez3_min_max.go che legge da tastiera
una serie di numeri (n numeri? o != 0 fino a incontrare 0)
e stampa il massimo e il minimo
*/

package main

import (
    "fmt"
)

func main() {
    var n,min,max int

    fmt.Print("Inserisci n: ")
    fmt.Scan(&n)

    for i := 0; i < n; i++{
      temp := 0
      fmt.Print(i,"\t- Inserisci un numero: ")
      fmt.Scan(&temp)

      if i > 1 {
        if temp > max {
          max = temp
        }else if temp < min {
          min = temp
        }
      }else if i == 0{
        min = temp
      }else if i == 1{
        max = temp
        if min > max{
          max = min
          min = temp
        }
      }
    }

    fmt.Println("Min: ",min," - Max: ",max)
}
