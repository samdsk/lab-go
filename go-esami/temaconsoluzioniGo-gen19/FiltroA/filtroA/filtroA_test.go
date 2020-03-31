/*
Filtro A
--------

Scrivere un programma filtroA.go che legge una sequenza di valori interi da linea di comando e controlla che si alternino valori pari e valori dispari.
In questo caso il programma stampa il messaggio "sequenza valida", altrimenti "elemento in posizione i non valido", dove i è la posizione del primo  elemento (da sinistra) che non rispetta la regola di alternanza o che non è un valore numerico.
In caso di mancanza di valori, il programma deve stampare "nessun valore in ingresso".
La sequenza può iniziare sia con un valore pari sia con uno dispari.
Si ricorda che lo zero è un numero pari.


Esempi
------

Input: 3 8 1 12
Output: sequenza valida

Input: 4
Output: sequenza valida

Input: 1 2 3 5
Output: elemento in posizione 4 non valido

Input: 1 2 3eqeqw 5
Output: elemento in posizione 3 non valido

Input:
Output: nessun valore in ingresso
*/
package main

import (
	"fmt"
	"log"
	"os/exec"

	"strings"
	"testing"
)

func TestNoParm(t *testing.T) {
	subproc := exec.Command("./filtroA") // presuppone compilato
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}

	fmt.Printf("Command line:\n%+s\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	if strings.TrimSpace(string(out)) != "nessun valore in ingresso" {
		t.Fail()
	}
}

func TestOneParm(t *testing.T) {
	subproc := exec.Command("./filtroA", "2") // presuppone compilato
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}

	fmt.Printf("Command line:\n%+s\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	if strings.TrimSpace(string(out)) != "sequenza valida" {
		t.Fail()
	}
}

func TestWrong1(t *testing.T) {
	subproc := exec.Command("./filtroA", "3", "8888", "0", "11", "57") // presuppone compilato
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}

	fmt.Printf("Command line:\n%+s\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	if strings.TrimSpace(string(out)) != "elemento in posizione 3 non valido" {
		t.Fail()
	}
}

func TestWrong2(t *testing.T) {
	subproc := exec.Command("./filtroA", "1", "2", "3", "5") // presuppone compilato
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}

	fmt.Printf("Command line:\n%+s\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	if strings.TrimSpace(string(out)) != "elemento in posizione 4 non valido" {
		t.Fail()
	}
}

func TestWrong3(t *testing.T) {
	subproc := exec.Command("./filtroA", "112", "1", "2", "3", "5") // presuppone compilato
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}

	fmt.Printf("Command line:\n%+s\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	if strings.TrimSpace(string(out)) != "elemento in posizione 5 non valido" {
		t.Fail()
	}
}

func TestWrong4(t *testing.T) {
	subproc := exec.Command("./filtroA", "5", "5", "10", "21", "40") // presuppone compilato
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}

	fmt.Printf("Command line:\n%+s\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	if strings.TrimSpace(string(out)) != "elemento in posizione 2 non valido" {
		t.Fail()
	}
}

func TestOk1(t *testing.T) {
	subproc := exec.Command("./filtroA", "4", "5", "12", "11", "0", "77") // presuppone compilato
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}

	fmt.Printf("Command line:\n%+s\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	if strings.TrimSpace(string(out)) != "sequenza valida" {
		t.Fail()
	}
}

func TestOk2(t *testing.T) {
	subproc := exec.Command("./filtroA", "3", "4", "5", "12", "11", "0", "77") // presuppone compilato
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}

	fmt.Printf("Command line:\n%+s\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	if strings.TrimSpace(string(out)) != "sequenza valida" {
		t.Fail()
	}
}

func TestWrongType1(t *testing.T) {
	subproc := exec.Command("./filtroA", "3", "4728", "5555y555", "-12", "-11", "0", "77") // presuppone compilato
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}

	fmt.Printf("Command line:\n%+s\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	if strings.TrimSpace(string(out)) != "elemento in posizione 3 non valido" {
		t.Fail()
	}
}

func TestWrongType2(t *testing.T) {
	subproc := exec.Command("./filtroA", "sadas3", "4728", "5555y555", "-12", "-11", "0", "77") // presuppone compilato
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}

	fmt.Printf("Command line:\n%+s\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	if strings.TrimSpace(string(out)) != "elemento in posizione 1 non valido" {
		t.Fail()
	}
}
