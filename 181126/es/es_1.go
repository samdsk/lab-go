/*
Creare un programma che gestisca uno stack generico (numero arbitrario di posizioni)

Lo stack è un array di valori (float) che mantiene nella prima posizione il valore più recente.
E' possibile accedere sempre e solo alla prima posizione.
Operazioni disponibili sullo stack:
- push -> aggiunge un valore in testa allo stack
- pop -> rimuove il valore in testa allo stack e lo ritorna
- top (o peek) -> ritorna il valore in testa allo stack senza rimuoverlo
- eventualmente empty -> lo stack è vuoto vero o falso?

ESEMPI:
stack -> [1, 5, 7, 1]
push(13)
stack -> [13, 1, 5, 7, 1]
n = pop()
n -> 13
stack -> [1, 5, 7, 1]
n = top()
n -> 1
stack -> [1, 5, 7, 1]

stack.go | tee filename.log //salva su .log quello che fa il prog
diff filename1 fliename2 
meld

*/

package main

import (
  "fmt"
  "strings"
)
var stack []float64

func main() {
  push(1)
  push(7)
  push(5)
  push(1)
  push(13)
  fmt.Println(pop(),":",stack)
  fmt.Println(top())
}

func push(f float64){
  stack = append([]float64{f},stack...)
}

func pop()float64{
  f := stack[0]
  stack = stack[1:]
  return f
}

func top()float64{
  return stack[0]
}

func str() string{
	return strings.Replace(fmt.Sprintf("%v",stack)," ", ", ", -1)
}


