/*
Parole
========

Scrivere un programma in Go (il file deve chiamarsi 'parole.go') che,
dato un file di input (il cui nome è dato su linea di comando),
crea una slice con tutte le "parole" (sequenze di caratteri alfabetici maiuscoli e minuscoli
trovate nel file stesso, usando quindi come separatori tutti i caratteri che non siano lettere,
cioè i caratteri di spaziatura, la punteggiatura, i numeri, i simboli, ecc.)
ed emette su standard output la slice stessa e il numero di "parole" trovate.

Suggerimento
------------
L'espressione regolare`[a-zA-Z]+` (o anche `\p{Latin}+`) definisce sequenze
di una o più lettere (senza differenziare tra minuscole e maiuscole).

Esempi
======

1) dato un file "uno.input" contenente:
sakdjhas872183612 2213821ouo pippo pippo123
pooiu23232ouoii        trtrtrt7676767ytyty

lanciando "go run stringhe.go uno.input" viene prodotto il seguente output:
[sakdjhas ouo pippo pippo pooiu ouoii trtrtrt ytyty]
8



2) dato un file "due.input" contenente:
WqeiQwuewq. 232wqeuqwe $uiUiU 98989::
++++8278 732873  "xnbcxnbcx[[]]000PPP"

lanciando "go run stringhe.go uno.input" viene prodotto il seguente output:
[WqeiQwuewq wqeuqwe uiUiU xnbcxnbcx PPP]
5

*/
package main

import (
	"fmt"
	"os/exec"

	//"strings"
	//"log"
	"testing"
)

func TestInput(t *testing.T) {
	lancia("uno.input", "[sakdjhas ouo pippo pippo pooiu ouoii trtrtrt ytyty]\n8\n", t)
	lancia("due.input", "[WqeiQwuewq wqeuqwe uiUiU xnbcxnbcx PPP]\n5\n", t)
}

func lancia(input string, output string, t *testing.T) {
	fmt.Println("### Questo test verifica output atteso!")
	subproc := exec.Command("./parole", input) // presuppone che sia già stato compilato (dallo script)

	//in, err := os.Open(input)
	//subproc.Stdin = in
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
