/**
Esercizio: calcolatrice RPN (il file deve chiamarsi RPN.go)

https://it.wikipedia.org/wiki/Notazione_polacca_inversa

per vederne una in funzione: http://www.alcula.com/calculators/rpn/

versione semplice:
- 2 accumulatori
- 4 operazioni (+ - * /)
- valori float

versione full: array di accumulatori


Implementazione LOOP:

- legge input utente
 - se numero accumula
 - se operatore collassa

- condizione di uscita (inserimento operatore "q")
- condizioni di errore


A1[]
A2[]

sequenza 3 4 5 6 +

risultato (per iterazioni):

A1[3]
A2[]

A1[4]
A2[3]

A1[5]
A2[4]

A1[6]
A2[5]

A1[11]
A2[]


IMPORTANTE: quando finito (a qualunque livello decidiate di mettervi, se cioè con 2 accum o array) caricate su upload.di.unimi.it che poi facciamo girare i test

*/
package main

import (
  "fmt"
  "strconv"
)


var stack []float64

func main() {

  var s string
  fmt.Println("Calcolatrice RPN")

  for{
    fmt.Print("Inserisci un numero or un operatore (q = exit): ")
    fmt.Scan(&s)

    f,err := strconv.ParseFloat(s,64)

    switch {
      case err == nil:
        push(f)

      case s == "+":
        if len(stack)<=1{
          fmt.Println("Errore: non hai inserito almeno due numeri!")
          break
        }
        b := stack[0]
        a := stack[1]
        pop()
        push(a+b)

      case s == "-":
        if len(stack)<=1{
          fmt.Println("Errore: non hai inserito almeno due numeri!")
          break
        }
        b := stack[0]
        a := stack[1]
        pop()
        push(a-b)

      case s == "*":
        if len(stack)<=1{
          fmt.Println("Errore: non hai inserito almeno due numeri!")
          break
        }
        b := stack[0]
        a := stack[1]
        pop()
        push(a*b)

      case s == "/":
        if len(stack)<=1{
          fmt.Println("Errore: non hai inserito almeno due numeri!")
          break
        }
        b := stack[0]
        a := stack[1]
        pop()
        push(a/b)

      case s == "q":
        fmt.Println("Result:",stack[0])
        fmt.Println("Exiting...")
        return
      default :
        fmt.Println("Errore: hai inserito un carattere/i non valido")
    }
  }

}

func push(f float64){
  stack = append([]float64{f},stack...)
}

func pop(){
  stack = stack[1:]
}
