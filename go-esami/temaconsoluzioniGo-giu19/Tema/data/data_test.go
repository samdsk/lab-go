/**

Esercizio "data"
================

Scrivere un programma Go (il file deve chiamarsi 'data.go') che:

- legge da linea di comando una stringa del tipo

	"2019-06-05 14:11:25"

  NB: anno-mese-giorno ore:minuti:secondi, *in un solo argomento tra virgolette*

- determina se i formati della data e dell'ora sono validi:
	- mese tra 01 e 12, giorno tra 01 e 31, anno tra 1900 e 2100, con separatore '-'
	  e giorno e mese di due caratteri
	- ore, minuti e secondi di due caratteri, con separatore ':'

- determina se l'ora è valida (ora tra 00 e 23, minuti e secondi tra 00 e 59)

- se i formati della data e dell'ora sono validi e l'ora è valida, stampa la data nel formato

	gg mese aaaa

con il nome del mese in italiano; altrimenti stampa il messaggio "argomento non valido"


Esempi
------

$ go run data.go "2019-06-05 12:11:25"
05 giugno 2019

$ go run data.go "2019-06-55 12:11:25"
argomento non valido

$ go run data.go "2019-06-05 12-11-25"
argomento non valido

$ go run data.go "2019-11-30 12:11:25"
30 novembre 2019

$ go run data.go "2019-11-30 12:11:75"
argomento non valido

$ go run data.go "2019-13-31 12:11:25"
argomento non valido

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
	lanciaEcontrolla("2019-06-05 12:11:25", "05 giugno 2019\n", t)
	lanciaEcontrolla("2019-6-5 12:11:25", "argomento non valido\n", t)
	lanciaEcontrolla("2019-11-30 12:11:25", "30 novembre 2019\n", t)
	lanciaEcontrolla("2019-11-30 12:11:75", "argomento non valido\n", t)
	lanciaEcontrolla("2019-06-55 12:11:25", "argomento non valido\n", t)
	lanciaEcontrolla("2019-06-05 12-11-25", "argomento non valido\n", t)
	lanciaEcontrolla("2019-13-31 12:11:25", "argomento non valido\n", t)
	lanciaEcontrolla("2019-10-30 12:11:25", "30 ottobre 2019\n", t)
	lanciaEcontrolla("2019-09-03 12:11:25", "03 settembre 2019\n", t)
	lanciaEcontrolla("2219-11-31 12:11:25", "argomento non valido\n", t)
	lanciaEcontrolla("2019-10-3o 12:11:25", "argomento non valido\n", t)
}

func lanciaEcontrolla(in string, out string, t *testing.T) {
	//fmt.Println(runtime.Caller(1));
	fmt.Println("-----------------------------------")
	subproc := exec.Command("./data", in) // presuppone che sia già stato compilato
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
