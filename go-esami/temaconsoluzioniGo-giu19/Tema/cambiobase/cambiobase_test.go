/**

Esercizio "cambio base"
=======================

Scrivere un programma Go (il file deve chiamarsi 'cambiobase.go') dotato di una funzione

- func tenToB(n, b int) (string, err)

che restituisce la rappresentazione in base b (b è il secondo parametro)
di un numero n (intero non negativo, in base 10) passato come primo parametro
e un valore di errore, diverso da nil se la base b non è compresa tra 2 e 16;

E di una funzione

- func main()

che legge due numeri interi, n e b, da linea di comando, ed emette
su standard output n espresso in base b.

Se il numero di parametri è diverso da due, se i parametri non sono
numeri interi, o se la base non è tra 2 e 16,
il programma deve stampare il messaggio corrispondente:
- "numero parametri non valido"
- "parametro non valido"
- "base non valida"

Suggerimento. Le cifre del risultato della conversione si ottengono prendendo
i resti della divisione di n per b, dividendo n per b fino a che il risultato
della divisione diventa 0.
I resti così ottenuti, in ordine inverso, formeranno il numero convertito.

Consideriamo ad esempio 6 in base 2

resto divisione 6 per 2 -> 0
quoziente divisione 6 per 2 -> 3
resto divisione 3 per 2 -> 1
quoziente divisione 3 per 2 -> 1
resto divisione 1 per 2 -> 1
quoziente divisione 1 per 2 -> 0 (stop)

Prendendo i resti in ordine inverso si ottiene il risultato: 110

Attenzione: per rappresentare numeri in base > 10 (e fino a base 16)
si devono usare anche le lettere A,B,C,D,E,F.
A tale scopo può essere conveniente usare la stringa "0123456789ABCDEF".

Esempi
------

$ go run cambiobase.go 13 3
111

$ go run cambiobase.go 4 20
base non valida

$ go run cambiobase.go 123 7
234

$ go run cambiobase.go 785 iuy
parametro non valido

$ go run cambiobase.go 140 13
AA

$ go run cambiobase.go 171 16
AB

$ go run cambiobase.go 123 2 4
numero parametri non valido

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
	lanciaEcontrolla("5", "2", "101\n", t)
	lanciaEcontrolla("7", "9", "7\n", t)
	lanciaEcontrolla("13", "10", "13\n", t)
	lanciaEcontrolla("13", "3", "111\n", t)
	lanciaEcontrolla("123", "7", "234\n", t)
	lanciaEcontrolla("140", "13", "AA\n", t)
	lanciaEcontrolla("255", "16", "FF\n", t)
	lanciaEcontrolla("785", "11", "654\n", t)
	lanciaEcontrolla("4", "20", "base non valida\n", t)
	lanciaEcontrolla("4.777", "3", "parametro non valido\n", t)
	lanciaEcontrolla("456", "tre", "parametro non valido\n", t)
}

func lanciaEcontrolla(in1 string, in2 string, out string, t *testing.T) {
	fmt.Println("-----------------------------------")
	subproc := exec.Command("./cambiobase", in1, in2) // presuppone che sia già stato compilato
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

// giusto per testare se la funzione c'è
func TestTenToB(t *testing.T) {
	fmt.Println("-----------------------------------")
	n_b, err := tenToB(5, 2)
	if err != nil {
		t.Errorf("Fallito: %s\n", err)
	}
	if string(n_b) != "101" {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
	fmt.Println(string(n_b))
}

func TestTenToBerror(t *testing.T) {
	fmt.Println("-----------------------------------")
	_, err := tenToB(5, 20)
	if err == nil {
		t.Errorf("Fallito: %s\n", err)
	}
}
