/*
Statistiche
===========

Scrivere un programma in Go (il file deve chiamarsi 'statistiche.go') dotato di

- una funzione
	moda(serie []float64) float64
  che, data una serie di numeri, restituisce la loro moda (*). In presenza di due (o più) mode,
  si restituisca la più piccola.

- una funzione
  	mediana (serie []float64) float64
  che, data una serie di numeri, restituisce la loro mediana (**)

- una funzione
	main()
  che legge una serie di numeri float64 da standard input, uno per riga,
  (la lettura viene terminata con invio seguito da Ctrl D))
  ed emette su standard output la loro moda e la loro mediana.


Note
----

Sia la moda che la mediana sono indipendenti dall'ordine dei valori nella serie.

(*) La moda di una serie di numeri è il valore che compare più frequentemente nella serie.

(**) La mediana di una serie di numeri è
	- nel caso di serie di lunghezza dispari,
	  il valore centrale della serie ordinata in ordine crescente
	- nel caso di serie di lunghezza pari, la media dei due valori centrali
	  della serie ordinata in ordine crescente

Suggerimenti
------------
Il package sort fornisce funzioni utili per ordinare una serie di valori

Esempi
======

1) data la serie in input:
1
2
5
4
3

si ottiene:

moda: 1
mediana: 3


2) data la serie in input:
4
3
2
1

si ottiene:

moda: 1
mediana: 2.5


3) data la serie in input:
4.1
2.7
4.1
2.7
2.7
4.1
2
5

si ottiene:

moda: 2.7
mediana: 3.4

*/
package main

import (
	"fmt"
	"os"
	"os/exec"

	//"strings"
	//"log"
	"testing"
)

func TestInput(t *testing.T) {
	lancia("uno.input", "moda: 1\nmediana: 3\n", t)
	lancia("due.input", "moda: 1\nmediana: 2.5\n", t)
	lancia("tre.input", "moda: 2.7\nmediana: 3.4\n", t)
}

func lancia(input string, output string, t *testing.T) {
	fmt.Println("### Questo test verifica output atteso!")
	subproc := exec.Command("./statistiche") // presuppone che sia già stato compilato (dallo script)

	in, err := os.Open(input)
	subproc.Stdin = in
	out, err := subproc.CombinedOutput()

	if err != nil {
		t.Errorf("Fallito: %s\n", err)
	}

	if output != string(out) {
		t.Errorf("Output NON corrisponde all'atteso")
	}

	fmt.Printf("Output ATTESO:\n%s\n", output)
	fmt.Printf("Output OTTENUTO:\n%s\n", string(out))

	subproc.Process.Kill()
	//subproc.Wait()
}

func TestModa(t *testing.T) {
	fmt.Println("-----moda------------------------------")

	s := []float64{2.7, 2.7, 2.7, 4.1, 4.1, 4.1, 2, 5}

	if 2.7 != moda(s) {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
}

func TestMediana(t *testing.T) {
	fmt.Println("-----mediana------------------------------")

	s := []float64{2.7, 2.7, 2.7, 4.1, 4.1, 4.1, 2, 5}

	if 3.4 != mediana(s) {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
}
