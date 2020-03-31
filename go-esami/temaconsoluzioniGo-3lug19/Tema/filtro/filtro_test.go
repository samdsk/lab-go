/*

Filtro
======

Scrivere un programma in Go (il file deve chiamarsi 'filtro.go') che traccia l'andamento
di una serie di numeri interi letti da standard input, uno per riga.
In particolare il programma legge una sequenza di lunghezza arbitraria
(lunghezza > 1, terminare l'input con invio seguito da Ctrl-D)
di valori interi e, a partire dal secondo valore, produce in output una stringa di
'+', '=', '-' in corrispondenza di andamento positivo, costante o negativo rispettivamente,
cioè a seconda che il nuovo valore letto sia maggiore, uguale o minore del valore precedente.



Esempi
======

1) data in input la sequenza:
34
30
25
34
37
37
37
37

il programma genera la sequenza:
--++===



2) data in input la sequenza:
-938719
0
-736273
10
12
13
13
11

il programma genera la sequenza:
+-+++=-

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
	lancia("uno.input", "--++===\n", t)
	lancia("due.input", "+-+++=-\n", t)
	lancia("tre.input", "+++++=----\n", t)
}

func lancia(input string, output string, t *testing.T) {
	fmt.Println("### Questo test verifica output atteso!")
	subproc := exec.Command("./filtro") // presuppone che sia già stato compilato (dallo script)

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
