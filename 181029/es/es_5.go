/*
5. Scrivere un programma lez3_primo.go che,
dato un intero, determina se un numero e` primo.
*/

package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Print("inserisci il numero da controllare: ")
    fmt.Scan(&n)

    var check bool
    if (n == 2){
      fmt.Println(n,"è un numero primo.")
      return
    }
    for i := 2; i<n; i++{
      temp := n%i

      if (temp == 0){
        //fmt.Println("Breaking point is i=",i)
        check = false
        break
      }else{
        check = true
      }
    }

    if check{
      fmt.Println(n,"è un numero primo.")
    }else{
      fmt.Println(n,"NON è un numero primo.")
    }
}
