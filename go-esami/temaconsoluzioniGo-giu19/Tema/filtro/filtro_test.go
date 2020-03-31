/**

Esercizio filtro
================

Scrivere un programma in Go (il file deve chiamarsi 'filtro.go') che,
data una serie (di lunghezza arbitraria) di numeri interi da standard input
(la serie contiene almeno un elemento, e termina con CTRL-D su una riga vuota),
calcola ed emette su standard output:

- il valore minimo (int)

- il valore massimo (int)

- la media (in tipo float64, quindi non troncata)

dei valori letti. Per il formato si vedano gli esempi qui sotto.

Si fornisce un file "esempioLungo.input" d'esempio, a fronte del quale l'output *dovrà*
essere il seguente:

min: -3556
max: 9292
media: 2871.39393939394

E un file "esempioBreve.input" d'esempio, a fronte del quale l'output *dovrà*
essere il seguente:

min: -26
max: 77
media: 27.75

Nota bene: la verifica della correttezza del filtro avviene mediante controllo automatico
(uno script) ergo attenersi strettamente all'output specificato sopra,
pena una possibile esclusione dalla prova d'esame.

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
	lancia("esempioLungo.input", t)
	lancia("esempioBreve.input", t)
	lancia("esempioPos.input", t)
	lancia("esempioNeg.input", t)
}

func lancia(input string, t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./filtro") // presuppone che sia già stato compilato (dallo script)

	in, err := os.Open(input)
	subproc.Stdin = in
	out, err := subproc.CombinedOutput()

	if err != nil {
		t.Errorf("Fallito: %s\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(out))

	subproc.Process.Kill()
	//subproc.Wait()
}
