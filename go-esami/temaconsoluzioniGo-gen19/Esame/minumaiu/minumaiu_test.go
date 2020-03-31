/*
MinuMaiu
--------

Scrivere un programma minumaiu.go, dotato di:
-  una funzione
	estraiParole(testo, r_exp string) (parole []string)
che estrae da testo le (sotto)stringhe che corrispondono a r_exp
- una funzione
	main()
che legge da un file, il cui nome è passato su linea di comando,
e, per ogni riga del file, stampa quante stringhe contengono
pattern <alternanza di mauiscola-minuscola>.
Alla fine stampa tutte le parole trovate.

Per chiarire, la stringa AbBc e` tutta <alternanza di mauiscola-minuscola>,
ma anche la stringa eSaMelab contiene il pattern
<alternanza di mauiscola-minuscola>.

Esempio
-------

Input (da file):
AbBc

 	casa mia
Ah MmMmMm
oHoH abcdEfGhilmnOp

Output:
riga 1 : 1 parola/e con il pattern
riga 2 : 0 parola/e con il pattern
riga 3 : 0 parola/e con il pattern
riga 4 : 2 parola/e con il pattern
riga 5 : 2 parola/e con il pattern
stringhe trovate : [AbBc Ah MmMmMm Ho EfGh Op]
*/
package main

// TODO aggiornare rispetto a nuova versione minumaio

import (
	"fmt"
	"os/exec"

	//"strings"
	"log"
	"testing"
)

func TestPatternAAA(t *testing.T) {
	if estraiParole("uno due aaaaaaaaa quattro", "[a]+")[0] != "aaaaaaaaa" {
		t.Fail()
	}
}

func TestPatternMinuMaiu(t *testing.T) {
	if estraiParole("uno due tre aBgThY cinque aBaB sei", "([[:lower:]][[:upper:]])+")[1] != "aBaB" {
		t.Fail()
	}
}

func TestIsPatternNumbers(t *testing.T) {
	if estraiParole("a b c 83768736638 u i 7676 o ppp 98", "[0-9]+")[2] != "98" {
		t.Fail()
	}
}

func TestInputSorgente(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./minumaiu", "./minumaiu_test.go") // presuppone che sia già stato compilato
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(out))
}
