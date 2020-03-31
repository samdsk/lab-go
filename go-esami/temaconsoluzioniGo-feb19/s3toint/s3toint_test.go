/*

Sistema di numerazione n-ario
-----------------------------

Si consideri il sistema di numerazione posizionale S3 in base 3 avente come
cifre i tre simboli
	z u d
i cui valori sono rispettivamente:
  val(z) = 0
  val(u) = 1
  val(d) = 2

Realizzare un programma s3toint.go che, dato un numerale num in S3 come argomento
sulla linea di comando, stampi il valore decimale di num.

Il programma deve (unicamente) controllare che l’argomento sia un numerale in S3 e
stampare “input non valido” nel caso non lo sia. Non sono richiesti altri controlli
(si può dare per scontato che l'argomento ci sia).

E` possibile dotare il programma di funzioni aggiuntive oltra a main().

Nota. Si ricorda che un sistema di numerazione si dice posizionale se i simboli
(cifre) usati per scrivere i numeri assumono valori diversi a seconda della
posizione che occupano. Il sistema di numerazione che usiamo solitamente è
un sistema posizionale in base 10.

Esempi
------
$ go run S3ToInt.go uzz
9
$ go run S3ToInt.go dddd
80
$ go run S3ToInt.go duz
21
$ go run S3ToInt.go uzc
input non valido
$ go run S3ToInt.go udzzzduuz
11001
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
	lanciaEcontrolla("z", "0\n", t)
	lanciaEcontrolla("uzz", "9\n", t)
	lanciaEcontrolla("dddd", "80\n", t)
	lanciaEcontrolla("duz", "21\n", t)
	//lanciaEcontrolla("temporaneo per verifica fail", "9\n", t)
	lanciaEcontrolla("uzc", "input non valido\n", t)
	lanciaEcontrolla("udzzzduuz", "11001\n", t)
	lanciaEcontrolla("akjhdaksjhsakdh", "input non valido\n", t)
}

func lanciaEcontrolla(in string, out string, t *testing.T) {
	fmt.Println("-----------------------------------")
	subproc := exec.Command("./s3toint", in) // presuppone che sia già stato compilato
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
