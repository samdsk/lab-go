/*
Triangolo
---------

Scrivere un programma triangolo.go che verifica l'appartenenza
(cioè all'interno o sul suo perimetro) di un punto (x,y) a un triangolo ABC,
la cui base AB si trova sull'asse delle ascisse e di cui sono date
le equazioni delle rette su cui stanno i lati AC e BC (vedi figura).


       r2\  /r1
   y|     \/
    |   C /\  .pE
    |    /  \
    |   / .pI\
    |__/______\________
	   A       B     x

Nota: Una retta è specificata attraverso il coefficiente angolare m
e il termine noto (intercetta) q della sua equazione y = m·x+q.

In particolare il programma stampa il punto e uno dei messaggi
(vedi anche esempi sotto)

- fuori
- dentro

a seconda che il punto appartenga o no.

Il programma deve essere dotato di

- una struttura Punto con due campi float64
per le coordinate x e y del punto

- una struttura Retta con due campi float64
per il coefficiente angolare m e l'intercetta q della retta

- una funzione
	Sotto(p Punto, r Retta) bool
che restituisce 'true' se il Punto p è sotto o sulla Retta r,
'false' altrimenti

- una funzione main()
che legge da standard input, nell'ordine,
	la m e la q della prima retta
	la m e la q della seconda retta
	la x e la y di un punto
verifica se il punto appartiene o no
al triangolo e stampa il messaggio opportuno.

Non si devono fare controlli sulle rette, si assuma cioè che definiscano un triangolo,
come mostrato nella figura sopra.

ESEMPI

go run triangolo.go
2 -2
-2 20
3 5
{3 5} fuori

go run triangolo.go
1 4
-1 12
1 1
{1 1} dentro

*/
package main

import (
	"fmt"
	"os/exec"
	"strings"

	//"log"
	"testing"
)

func TestMainMultiplo(t *testing.T) {
	lanciaEcontrolla("1 0 -1 4 .2 .1", "{0.2 0.1} dentro\n", t)
	lanciaEcontrolla("1 0 -1 5 1 .5", "{1 0.5} dentro\n", t)
	lanciaEcontrolla("1 0 -1 10 3 2", "{3 2} dentro\n", t)
	lanciaEcontrolla("1 0 -1 10 13 2", "{13 2} fuori\n", t)
}

func TestMainSottoAscisse(t *testing.T) {
	lanciaEcontrolla("1 0 -1 4 .2 -.1", "{0.2 -0.1} fuori\n", t)
	lanciaEcontrolla("1 0 -1 5 1 -.5", "{1 -0.5} fuori\n", t)
	lanciaEcontrolla("1 0 -1 10 3 -2", "{3 -2} fuori\n", t)
}

func lanciaEcontrolla(in string, out string, t *testing.T) {
	fmt.Println("-----------------------------------")
	subproc := exec.Command("./triangolo", in) // presuppone che sia già stato compilato
	subproc.Stdin = strings.NewReader(in)
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

func TestSotto1(t *testing.T) {
	p := Punto{1, 1}
	r := Retta{1, 10}
	if !Sotto(p, r) {
		t.Fail()
	}
}

func TestSotto2(t *testing.T) {
	p := Punto{11, 11}
	r := Retta{0, 10}
	if Sotto(p, r) {
		t.Fail()
	}
}

func TestSotto3(t *testing.T) {
	p := Punto{3, 3}
	r := Retta{1, 0}
	if !Sotto(p, r) {
		t.Fail()
	}
}
