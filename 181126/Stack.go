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
*/
package main

import (
    "fmt"
    "strings"
)

var stack []float32

func main() {
    push(1)
    push(7)
    push(5)
    push(1)
    fmt.Println("stack ->", str())
    push(13)
    fmt.Println("stack ->", str())
    fmt.Println("n ->", pop())
    fmt.Println("stack ->", str())
    fmt.Println("n ->", top())
    fmt.Println("stack ->", str())
}

func pop() float32 {
    var r float32
    r, stack = stack[0], stack[1:]
    return r
}

func top() float32 {
	return stack[0]
}

func push(value float32) {
	stack = append([]float32{value}, stack...)
}

func empty() bool {
	return len(stack) == 0
}

func str() string {
    return strings.Replace(fmt.Sprintf("%v", stack), " ", ", ", -1)
}
