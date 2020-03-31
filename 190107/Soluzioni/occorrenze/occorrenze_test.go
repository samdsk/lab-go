package main

import (
	"fmt"
	"log"
	"os/exec"
	//"strings"
	"testing"
)

var occTests = []struct {
	s        string // input
	r        rune   // input
	expected int    // expected result
}{
	{"pippo", 'p', 3},
	{"pippo", 'i', 1},
	{"paparino", 'a', 2},
	{"metempsicosi", 'm', 2},
}

func TestOccorrenze(t *testing.T) {
	for _, tt := range occTests {
		actual := occorrenze(tt.s, tt.r)
		if actual != tt.expected {
			t.Errorf("occorrenze(%s,%c): expected %d, actual %d", tt.s, tt.r, tt.expected, actual)
		}
	}
}

func TestInput(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./occorrenze", "u", "aiuola grande come una casa") // presuppone che sia già stato compilato (dallo script)
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}
	fmt.Printf("Command line:\n%+q\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	subproc.Wait()
}
