/*
Frazioni
========

Scrivere un programma in Go (il file deve chiamarsi 'frazioni.go') dotato di:

- una struttura
	Fraction (num,den int)
  che rappresenta una frazione num/den, cioè dove num è il numeratore e den il denominatore.

- una funzione
	newFraction(n,d int) (Fraction, error)
  che, dati due interi, n e d, restituisce la frazione ridotta ai minimi termini (*)
  equivalente a n/d e un errore "divisore nullo" nel caso d sia uguale a 0.

- una funzione
	sum(f1, f2 Fraction) Fraction
  che restituisce la somma di f1 e f2, ridotta ai minimi termini

- una funzione
	main()
  che legge 4 valori interi non negativi da linea di comando (n1, d1, n2, d2),
  crea le due frazioni n1/d1 e n2/d2 ed emette su standard output:

  - il messaggio di errore "numero di parametri non valido" se il numero di parametri interi è diverso da 4

	- un messaggio di errore se una o tutte e due le frazioni hanno denominatore nullo

	- altrimenti la loro somma

	Si può assumere che i parametri passati da linea di comando siano tutti numeri di tipo int e non negativi.

(*) Suggerimento: Un modo (non efficiente ma semplice) di ridurre una frazione n/d ai
    minimi termini è dividere via via sia n sia d per ciascun divisore di entrambi n e d.



Esempi
======

1) eseguendo:

$ go run frazioni.go 1 2 1 4

si ottiene:

3 / 4


2) eseguendo:

$ go run frazioni.go 4 12 2 0

si ottiene:

divisore nullo


3) eseguendo:

$ go run frazioni.go 3 10 19 5 7

si ottiene:

numero di parametri non valido


4) eseguendo:

$ go run frazioni.go 456 342 78 98

si ottiene:

313 / 147


5) eseguendo:

$ go run frazioni.go 4 12 2 3

si ottiene:

1 / 1

*/

package main

import (
	"fmt"
	"os/exec"

	//"strings"
	//"log"
	"testing"
)

func TestMainMultiplo(t *testing.T) {
	lanciaEcontrolla("1", "20", "1", "30", "1 / 12\n", t)
	lanciaEcontrolla("1", "2000", "1", "2000", "1 / 1000\n", t)
	lanciaEcontrolla("1", "2", "1", "4", "3 / 4\n", t)
	lanciaEcontrolla("4", "12", "2", "0", "divisore nullo\n", t)
	lanciaEcontrolla("456", "342", "78", "98", "313 / 147\n", t)
	lanciaEcontrolla("4", "12", "2", "3", "1 / 1\n", t)
}

func lanciaEcontrolla(n1, d1, n2, d2, out string, t *testing.T) {
	fmt.Println("----main-------------------------------")
	subproc := exec.Command("./frazioni", n1, d1, n2, d2) // presuppone che sia già stato compilato
	stdout, stderr := subproc.CombinedOutput()

	if stderr != nil {
		t.Errorf("Fallito: %s\n", stderr)
	}
	fmt.Printf("Output:\n%s\n", string(stdout))
	fmt.Printf("Expected output:\n%s\n", out)
	if string(stdout) != out {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
	subproc.Process.Kill()
}

func TestNumParSbagliato(t *testing.T) {
	fmt.Println("-----main par sbagliati------------------------------")
	subproc := exec.Command("./frazioni", "1", "2", "3", "4", "5") // presuppone che sia già stato compilato
	stdout, stderr := subproc.CombinedOutput()

	if stderr != nil {
		t.Errorf("Fallito: %s\n", stderr)
	}
	fmt.Printf("Output:\n%s\n", string(stdout))

	out := "numero di parametri non valido\n"
	fmt.Printf("Expected output:\n%s\n", out)
	if string(stdout) != out {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
	subproc.Process.Kill()
}

// giusto per testare se la funzione c'è
func TestNewFrac(t *testing.T) {
	fmt.Println("-----new frac------------------------------")

	fr, err := newFraction(5, 2)
	if err != nil {
		t.Errorf("Fallito: %s\n", err)
	}

	//fr2 := Frazione{5, 2}
	fr2 := Fraction{5, 2}

	if fr != fr2 {
		fmt.Println(">>> newFraction FAIL!")
		t.Fail()
	}
	fmt.Println(fr.num, fr.den)
}

func TestSum(t *testing.T) {
	fmt.Println("-----sum------------------------------")

	fr1, err1 := newFraction(2, 5)
	fr2, err2 := newFraction(3, 5)
	res, errR := newFraction(1, 1)
	if err1 != nil || err2 != nil || errR != nil {
		t.Errorf("Fallito")
	}

	if res != sum(fr1, fr2) {
		fmt.Println(">>> sum FAIL!")
		t.Fail()
	}
}
