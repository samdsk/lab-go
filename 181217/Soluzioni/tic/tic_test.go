package main

import (
	"fmt"
	//"log"
	//"strings"
	"os/exec"
	"testing"
)

func TestNewOrarioOK(t *testing.T) {
	esito, _ := newOrario(30, 30, 10)
	if !esito {
		t.Fail()
	}
}

func TestNewOrarioSBAGLIATO(t *testing.T) {
	esito, _ := newOrario(70, 30, 10)
	if esito {
		t.Fail()
	}
}

func TestTicOrarioNORMALE(t *testing.T) {
	_, orig := newOrario(30, 30, 10)
	_, target := newOrario(31, 30, 10)

	tic(&orig)

	if orig != target {
		t.Fail()
	}
}

func TestTicOrarioSCAVALLOMIN(t *testing.T) {
	_, orig := newOrario(59, 30, 10)
	_, target := newOrario(0, 31, 10)

	tic(&orig)

	if orig != target {
		t.Fail()
	}
}

func TestTicOrarioSCAVALLOH(t *testing.T) {
	_, orig := newOrario(59, 59, 10)
	_, target := newOrario(0, 0, 11)

	tic(&orig)

	if orig != target {
		t.Fail()
	}
}

func TestTicOrarioSCAVALLODAY(t *testing.T) {
	_, orig := newOrario(59, 59, 23)
	_, target := newOrario(0, 0, 0)

	tic(&orig)

	if orig != target {
		t.Fail()
	}
}

/*
func TestTic2(t *testing.T) {
}
*/

func TestMain1(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./tic", "65", "3", "2", "56789") // presuppone che sia già stato compilato (dallo script)
	out, err := subproc.CombinedOutput()

	if err != nil {
		//log.Fatalf("Fallito: %s\n", err)
		fmt.Printf("Fallito: %s\n", err)
	}
	fmt.Printf("Command line:\n%+q\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	subproc.Wait()
}

func TestMain2(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./tic", "5", "4", "3", "56789") // presuppone che sia già stato compilato (dallo script)
	out, err := subproc.CombinedOutput()

	if err != nil {
		//log.Fatalf("Fallito: %s\n", err)
		fmt.Printf("Fallito: %s\n", err)
	}
	fmt.Printf("Command line:\n%+q\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	subproc.Wait()
}

func TestMainNegativo(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./tic", "5", "-4", "3", "56789") // presuppone che sia già stato compilato (dallo script)
	out, err := subproc.CombinedOutput()

	if err != nil {
		//log.Fatalf("Fallito: %s\n", err)
		fmt.Printf("Fallito: %s\n", err)
	}
	fmt.Printf("Command line:\n%+q\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	subproc.Wait()
}
