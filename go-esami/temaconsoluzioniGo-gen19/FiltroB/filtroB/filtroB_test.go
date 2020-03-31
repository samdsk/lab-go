/*
Filtro B
--------

Scrivere un programma filtroB.go che legge una sequenza di valori interi da linea di comando e controlla che si alternino valori negativi e valori non negativi (positivi o zero).
In questo caso il programma stampa il messaggio "sequenza valida", altrimenti "elemento in posizione i non valido", dove i è la posizione del primo  elemento (da sinistra) che non rispetta la regola di alternanza o che non è un valore numerico.
In caso di mancanza di valori, il programma deve stampare "nessun valore in ingresso".
La sequenza può iniziare sia con un valore positivo che con uno negativo.


Esempi
------

Input: -11 0 -6 8
Output: sequenza valida

Input: 4
Output: sequenza valida

Input: 0 -4 2 23 -12
Output: elemento in posizione 4 non valido

Input: 0 -4 2 2erere3 -12
Output: elemento in posizione 4 non valido

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
	subproc := exec.Command("./filtroB") // presuppone compilato
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
	subproc := exec.Command("./filtroB", "2") // presuppone compilato
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

func TestOneParmNonNum(t *testing.T) {
	subproc := exec.Command("./filtroB", "2sdasd") // presuppone compilato
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

func TestOk1(t *testing.T) {
	subproc := exec.Command("./filtroB", "-1056", "0", "-6", "8") // presuppone compilato
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
	subproc := exec.Command("./filtroB", "23", "-1", "0", "-600", "8") // presuppone compilato
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
	subproc := exec.Command("./filtroB", "2", "-5", "-9898", "7676", "972") // presuppone compilato
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
	subproc := exec.Command("./filtroB", "-4", "2", "3", "0") // presuppone compilato
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

func TestOkCycle(t *testing.T) {
	subproc := exec.Command("./filtroB", "-467", "467", "-467", "467") // presuppone compilato
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

func TestWrong3(t *testing.T) {
	subproc := exec.Command("./filtroB", "3", "-4728", "5", "-12", "-11", "0", "77") // presuppone compilato
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

func TestWrongType1(t *testing.T) {
	subproc := exec.Command("./filtroB", "3", "-4728", "555k5555", "-12", "-11", "0", "77") // presuppone compilato
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
	subproc := exec.Command("./filtroB", "3sadkjash", "-4728", "555k5555", "-12", "-11", "0", "77") // presuppone compilato
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
