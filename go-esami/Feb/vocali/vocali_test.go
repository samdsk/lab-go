/*

Vocali
------

Scrivere un programma vocali.go che analizza un testo e conta
le occorrenze delle vocali (sia minuscole che maiuscole, ma non le accentate)
nel testo.
In particolare il programma è dotato di:

- una funzione per contare le occorrenze delle diverse vocali in una stringa
	func ContaVocali(s string, vocali map[rune]int)
che, data una stringa s e una mappa vocali, aggiorna opportunamente la mappa vocali
aggiungendo eventuali vocali non ancora presenti / incrementandone  i valori

- una funzione
	func main()
che legge un file, il cui nome è passato su linea di comando, produce
una mappa tra vocali presenti nel testo e il numero delle loro occorrenze nel testo,
e la stampa.
In mancanza di argomenti a linea di comando, il programma deve stampare:
	Mancano argomenti


ESEMPIO

se il file contiene:

	jdhkas c'è dkasjhkdjashkdh askdh ksah @@@ €€€ ### Ħ  wi) Ø
	qwqwe qwyewquteuqwte q 312312 2312wweqe €łłŧŧŧŧŧ sdasdas
	AA JKJLKLJLKJ LIIIIII
	U u ù
	aeiou
	AEIOU

l'output è (salvo l'ordine di stampa):
o : 1
E : 1
u : 4
U : 2
e : 7
A : 3
I : 7
O : 1
a : 8
i : 2

*/
package main

import (
	"fmt"
	"os/exec"

	//"strings"
	//"log"
	"testing"
)

func TestMainTestSource(t *testing.T) {
	fmt.Println(">>> ATTENZIONE! nessun controllo sulla validità dell'output, verificare a vista <<<")
	subproc := exec.Command("./vocali", "./vocali_test.go") // presuppone che sia già stato compilato
	out, err := subproc.CombinedOutput()

	if err != nil {
		t.Errorf("Fallito: %s\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(out))
	subproc.Process.Kill()
}

func TestMainNoArg(t *testing.T) {
	subproc := exec.Command("./vocali") // presuppone che sia già stato compilato
	out, err := subproc.CombinedOutput()

	if err != nil {
		t.Errorf("Fallito: %s\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(out))
	if string(out) != "Mancano argomenti\n" {
		t.Fail()
	}
	subproc.Process.Kill()
}

func TestContaVocali1(t *testing.T) {
	v := make(map[rune]int)
	ContaVocali("akdjhak sdhaskduaiq àèìòù oywioeoy ", v)
	if v['o'] != 3 || v['a'] != 4 || v['e'] != 1 || v['i'] != 2 || v['u'] != 1 {
		t.Fail()
	}
}

func TestContaVocali2(t *testing.T) {
	v := make(map[rune]int)
	ContaVocali("AIUOLE in FIORE", v)
	if v['A'] != 1 || v['U'] != 1 || v['i'] != 1 {
		t.Fail()
	}
}
