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

var stack []float32

func main() {
	/*TEST*/
}

func pop() float32 {
	/*TODO*/
}

func top() float32 {
	/*TODO*/
}

func push(value float32) {
	/*TODO*/
} 

func empty() bool {
	/*TODO*/
}
