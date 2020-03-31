/*
Percorso
--------

Scrivere un programma percorso.go che ha:
- una struttura Punto con due campi float64 x, y che
rappresentano le coordinate del punto nel piano cartesiano.

- una funzione
	distanza(p1,p2 Punto) float64
che, dati due Punti p1 e p2, restituisce la distanza tra p1 e p2.

- una funzione
	percorso(tappe []Punto) float64
che, data una sequenza di tappe, rappresentata da una slice di Punti,
restituisce la distanza da percorrere, partendo dalla prima tappa,
per tornare a questa, passando per tutte le altre tappe in ordine.

- una funzione
	main()
che legge da standard input una sequenza (non vuota) di valori numerici (floating point)
x1, y1, x2, y2, ..., xn, yn (per terminare la sequenza, dare invio e CTRL-D)
e stampa la sequenza di punti letti e la distanza totale da percorrere per andare
dal primo al secondo punto, dal secondo al terzo, ... e dall'ultimo al primo.

Nota : si puo` assumere, senza fare controlli, che la sequenza in input contenga
almento i dati di un punto e che i dati siano nella forma attesa (float).


Esempi
------

Input: 0 3 0 0 4 0
Output:
tappe: [{0 3} {0 0} {4 0}]
percorso 12

Input: 1.5 3.2 5.1 7.6
Output:
tappe: [{1.5 3.2} {5.1 7.6}]
percorso 11.370136322841516
*/
package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func TestDuePunti(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./percorso") // presuppone che sia già stato compilato (dallo script)
	input := "1 1 2 2"
	subproc.Stdin = strings.NewReader(input)
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}
	fmt.Printf("Input:\n%s\n", input)
	fmt.Printf("Output:\n%s\n", string(out))
}

func TestUnPunto(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./percorso") // presuppone che sia già stato compilato (dallo script)
	input := "1 1"
	subproc.Stdin = strings.NewReader(input)
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}
	fmt.Printf("Input:\n%s\n", input)
	fmt.Printf("Output:\n%s\n", string(out))
}

/*
func TestInvalidInput(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./percorso") // presuppone che sia già stato compilato (dallo script)
	input := "pippo pluto paperino"
	subproc.Stdin = strings.NewReader(input)
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}
	fmt.Printf("Input:\n%s\n", input)
	fmt.Printf("Output:\n%s\n", string(out))
}
*/

func TestPercorso(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./percorso") // presuppone che sia già stato compilato
	input := "5 6 -11 2 1.3 1 67.9 2 2 8"
	subproc.Stdin = strings.NewReader(input)
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}
	fmt.Printf("Input:\n%s\n", input)
	fmt.Printf("Output:\n%s\n", string(out))
}

func TestDispari(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./percorso") // presuppone che sia già stato compilato
	input := "1 1 4"
	subproc.Stdin = strings.NewReader(input)
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}
	fmt.Printf("Input:\n%s\n", input)
	fmt.Printf("Output:\n%s\n", string(out))
}
