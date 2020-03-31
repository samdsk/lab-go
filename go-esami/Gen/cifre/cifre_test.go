/*
Cifre
------

Scrivere un programma cifre.go dotato di
- una funzione
	contaCifre(s string) [10]int
che restituisce un array di lunghezza 10 che contiene il numero di
0, 1, 2, ... e 9 (in quest'ordine,) che si trovano nella stringa s.
Non ci sono vincoli sui tipi di caratteri in s.

- una funzione main() che legge una frase (tra "") da linea di comando che
contiene caratteri di qualsiasi tipo (cifre, lettere, simboli,
punteggiatura, lettere accentate, ecc.)
e stampa quante cifre di ogni tipo ci sono, ignorando gli eventuali altri caratteri.
Se manca l'argomento da linea di comando, deve stampare "manca argomento" e terminare.

Esempio
-------

Input: "1, 2, 3; non c’è fiaba senza re; 1, 2, 3; venite giù da me; 4, 5, 6; siete dei babbei; 7,8,9; io sono già altrove."
Output: [0 2 2 2 1 1 1 1 1 1]
*/
package main

import (
	"fmt"
	"os/exec"

	//"strings"
	"log"
	"testing"
)

func TestCompare1(t *testing.T) {
	if fmt.Sprintf("%v", contaCifre("1, 2, 3; non c’è fiaba senza re; 1, 2, 3; venite giù da me; 4, 5, 6; siete dei babbei; 7,8,9; io sono già altrove.")) != "[0 2 2 2 1 1 1 1 1 1]" {
		t.Fail()
	}
}

func TestCompare2(t *testing.T) {
	if fmt.Sprintf("%v", contaCifre("[0 2 2 2 1 1 1 1 1 1]")) != "[1 6 3 0 0 0 0 0 0 0]" {
		t.Fail()
	}
}

func TestCompare3(t *testing.T) {
	if fmt.Sprintf("%v", contaCifre("abcdefg")) != "[0 0 0 0 0 0 0 0 0 0]" {
		t.Fail()
	}
}

func TestCompare4(t *testing.T) {
	if fmt.Sprintf("%v", contaCifre("aaaaaaa1ddd3ddddd6gggg8hhhh")) != "[0 1 0 1 0 0 1 0 1 0]" {
		t.Fail()
	}
}

func TestMain(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./cifre", "abcde4fghi8askdjhaskdjash6sakdjhasdkjsah8") // presuppone che sia già stato compilato
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(out))
}

func TestMainNoArg(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./cifre") // presuppone che sia già stato compilato
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(out))
}

// verifica che sia proprio un array di int
func TestCheckType(t *testing.T) {
	var array [10]int
	array = contaCifre("1, 2, 3; non c’è fiaba senza re; 1, 2, 3; venite giù da me; 4, 5, 6; siete dei babbei; 7,8,9; io sono già altrove.")
	fmt.Println(array)
}
