/* Quoziente e resto

Problema Scrivere un programma Go che, dati un dividendo e un divisore
(interi), calcoli il quoziente e il resto.

Annotazioni L’operatore per la divisione (/) tra interi calcola la parte intera
del risultato. L’operatore per il resto della divisione è %.

Cosa succede se il divisore è 0?

e se sono entrambi 0?
*/

package main

import (
    "fmt"
)

func main() {
  var num1, num2 int64
  fmt.Println("Inserisci dividendo e un divisore")
  fmt.Scan(&num1,&num2)
  fmt.Println("Dividendo: ", num1, " Divisore: ", num2)

  fmt.Println("Quoziente: ", num1/num2)
  fmt.Println("Il resto: ", num1%num2)  
}
