/*
9. Scrivere un programma lez3_fibonacci.go che, dato un intero n,
stampa i primi n numeri della serie di Fibonacci: n_i = n_{i-1} + n_{i-2}
*/

package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Print("Inserisci un numero n>2: ")
    fmt.Scan(&n)

    for k:=1; k<=n; k++{
      numeratore,denominatore,k_1 := 0,0,0

      fmt.Println("k",k)
      n_check := 0

      if n-k <= 0{
        n_check = 1
      }else{
        n_check = n-k
      }

      for i:=0; n_check -i != 0;{
        numeratore =+ n_check-i
        i++
        fmt.Println("first")
      }

      d_check := 0

      if n-(2*k)+1 <= 0 {
        d_check = 1
      }else{
        d_check = n-(2*k)+1
      }

      for i:=0; d_check-i != 0;{
        denominatore =+ d_check-i
        i++
        fmt.Println("second",i)
      }

      for i:=0; (k-1-i) != 0;{
        k_1 =+ k-1-i
        i++
        fmt.Println("third")
      }

      if numeratore == 0 || n_check == 0 {
        numeratore = 1
      }
      if k_1  == 0{
        k_1 = 1
      }
      if denominatore == 0 || d_check == 0{
        denominatore = 1
      }

      temp := numeratore/(k_1*denominatore)

      mem := temp
      mem =+ temp
      fmt.Println("Serie di Fibonacci indice",k,":=",temp,"numeratore,deno,k_1:",numeratore,denominatore,k_1)
    }
}
