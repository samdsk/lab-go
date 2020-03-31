/*
Lessico
-------

Scrivere un programma lessico.go che
- stampa il seguente menu di opzioni
	+ (Legge una riga e memorizza le parole)
	? (Legge una parola e indica le righe che la contengono)
	p (Stampa le parole)

- legge stringhe da standard input
- il programma termina quando riceve un "end of file" (cioè EOF, pressione di 'CTRL-d')

Se la stringa inizia con:
- "+" (alimenta dizionario): il programma usa il rimanente della riga e memorizza in un "dizionario" le parole che la costituiscono;
- "?" (consulta dizionario): il programma usa il rimanente della riga e stampa i numeri di riga del dizionario 	in cui è comparsa la stringa;
- "p" (print): il programma stampa le parole presenti nel  "dizionario", con l’elenco dei numeri di riga in cui compaiono;


Esempio
-------

+ (Legge una riga e memorizza le parole)
? (Legge una parola e indica le riga che la contengono)
p (Stampa le parole)
+ la befana ha il fazzoletto e la gonna rattoppata
p
map[rattoppata:[1] la:[1 1] befana:[1] ha:[1] il:[1] fazzoletto:[1] e:[1] gonna:[1]]
+ ma quest'anno poverina la befana è raffreddata
? la
parola: la
righe [1 1 2]
? befana
parola: befana
righe [1 2]
? il
parola: il
righe [1]
p
map[raffreddata:[2] la:[1 1 2] il:[1] fazzoletto:[1] ma:[2] quest'anno:[2] poverina:[2] è:[2] befana:[1 2] ha:[1] e:[1] gonna:[1] rattoppata:[1]]

*/
package main

import (
	"fmt"
	"log"
	"os/exec"

	"strings"
	"testing"
	//"bytes"
	//"io"
	//"os"
)

func TestInput1(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./lessico") // presuppone che sia già stato compilato
	input := "+ aaa bbb ccc\np\n+ aaa bbb ccc\n+ uno due tre\np\n+ h j k\np"
	subproc.Stdin = strings.NewReader(input)
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}
	fmt.Printf("Input:\n%s\n", input)
	fmt.Printf("Output:\n%s\n", string(out))
}

func TestInput2(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./lessico") // presuppone che sia già stato compilato
	input := "+ la befana ha il fazzoletto e la gonna rattoppata\np\n+ ma quest'anno poverina la befana è raffreddata\n? la\n? befana\n? il\np"
	subproc.Stdin = strings.NewReader(input)
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}
	fmt.Printf("Input:\n%s\n", input)
	fmt.Printf("Output:\n%s\n", string(out))
}
