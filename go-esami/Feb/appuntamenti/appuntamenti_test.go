/*
Appuntamenti
------------

Scrivere un programma 'appuntamenti.go' dotato di:

- una struttura Appuntamento con tre campi int (giorno, oraInizio, oraFine)
che rappresentano il giorno dell'anno dell'appuntamento, l'ora di inizio
e quella di fine, rispettivamente.
Si considerano per semplicità solo appuntamenti che iniziano e finiscono
nella stessa giornata e a ore precise (intere).

- una funzione
	NewAppuntamento(gg, h1, h2 int) (Appuntamento, bool)
che, se i parametri sono validi (giorno tra 1 e 366,
ora inizio precedente o uguale a ora di fine e tutte e due tra 0 e 24)
crea un appuntamento con quei dati e lo restituisce insieme a 'true',
altrimenti restituisce un Appuntamento zero-value e 'false'.

- una funzione
    CheckSovrapposizioneAppuntamenti(app1, app2 Appuntamento) bool
che, dati due appuntamenti (si supponga siano validi), restituisce 'true'
se si sovrappongono (anche parzialmente), 'false' altrimenti.

- una funzione
	AddAppuntamento(app Appuntamento, agenda *[]Appuntamento) bool
che, se l'appuntamento app non si sovrappone ad altri già presenti in agenda,
lo aggiunge all'agenda e restituisce 'true', altrimenti restituisce 'false'

- una funzione
    AppuntamentiGiornata(gg int, agenda []Appuntamento) []Appuntamento
che, dati un giorno e una agenda di appuntamenti, estrae e restituisce gli
appuntamenti del giorno passato come parametro

- una funzione main() che legge da standard input
righe che contengono tre interi che rappresentanto nell'ordine

giorno ora-inizio ora-fine

(separatore è solo whitespace)

e per ciascuna aggiunge (se valido e non in sovrapposizione)
il corrispondente appuntamento.
La lettura termina con EOF (CTRL D) e a quel punto il programma
visualizza l'agenda degli appuntamenti (in ordine qualsiasi) su standard output.

*/
package main

import (
	//"strings"
	//"log"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestAppuntamentoNorm(t *testing.T) {
	_, correct := NewAppuntamento(300, 11, 12)
	if !correct {
		t.Fail()
	}
}

func TestAppuntamentoWrongDay(t *testing.T) {
	_, correct := NewAppuntamento(400, 11, 12)
	if correct {
		t.Fail()
	}
}

func TestAppuntamentoWrongTime(t *testing.T) {
	_, correct := NewAppuntamento(300, 11, 32)
	if correct {
		t.Fail()
	}
}

func TestSovrAppuntamentoPost(t *testing.T) {
	uno, _ := NewAppuntamento(300, 11, 13)
	due, _ := NewAppuntamento(300, 12, 14)
	if !CheckSovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
}

func TestSovrAppuntamentoFake(t *testing.T) {
	uno, _ := NewAppuntamento(302, 11, 13)
	due, _ := NewAppuntamento(300, 11, 13)
	if CheckSovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
}

func TestSovrAppuntamentoIdentici(t *testing.T) {
	uno, _ := NewAppuntamento(300, 11, 13)
	due, _ := NewAppuntamento(300, 11, 13)
	if !CheckSovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
}

func TestSovrAppuntamentoPre(t *testing.T) {
	uno, _ := NewAppuntamento(300, 11, 13)
	due, _ := NewAppuntamento(300, 10, 12)
	if !CheckSovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
}

func TestSovrAppuntamentoPreSameEnd(t *testing.T) {
	uno, _ := NewAppuntamento(300, 11, 13)
	due, _ := NewAppuntamento(300, 10, 13)
	if !CheckSovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
}

func TestSovrAppuntamentoFakeContigui(t *testing.T) {
	uno, _ := NewAppuntamento(300, 11, 13)
	due, _ := NewAppuntamento(300, 13, 15)
	if CheckSovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
}

func TestAppuntamentoReversedTime(t *testing.T) {
	_, correct := NewAppuntamento(300, 12, 10)
	if correct {
		t.Fail()
	}
}

func TestAgendaNorm(t *testing.T) {
	agenda := make([]Appuntamento, 0)
	for index := 300; index < 400; index += 20 {
		a, ok := NewAppuntamento(index, 11, 13)
		if ok {
			AddAppuntamento(a, &agenda)
			success := AddAppuntamento(a, &agenda)

			if success {
				t.Fail()
			}
		}
	}

	if len(agenda) != 4 {
		t.Fail()
	}

	if len(AppuntamentiGiornata(300, agenda)) != 1 {
		t.Fail()
	}
}

func TestMain1(t *testing.T) {
	subproc := exec.Command("./appuntamenti") // presuppone che sia già stato compilato
	input := "200 11 12\n201 11 12\n201 11 13\n201 13 12\n"
	subproc.Stdin = strings.NewReader(input)
	out, err := subproc.CombinedOutput()

	if err != nil {
		t.Errorf("Fallito: %s\n", err)
	}
	fmt.Printf("Input:\n%s\n", input)
	fmt.Printf("Output:\n%s\n", string(out))
	if string(out) != "[{200 11 12} {201 11 12}]\n" {
		t.Fail()
	}
	subproc.Process.Kill()
}
