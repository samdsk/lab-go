package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func TestInput1(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./filtro") // presuppone che sia già stato compilato (dallo script)
	input := "28 22 25 18 21 26 14 22 26 30 23 27 24 28 25 0"
	subproc.Stdin = strings.NewReader(input)
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}
	fmt.Printf("Input:\n%s\n", input)
	fmt.Printf("Output:\n%s\n", string(out))

	subproc.Wait()
}

func TestInput2(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./filtro") // presuppone che sia già stato compilato (dallo script)
	input := "28 30 18 22     3627	11\n25 26 18 21 26 14 22 26 30 23 27 24 28 25 0"
	subproc.Stdin = strings.NewReader(input)
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}
	fmt.Printf("Input:\n%s\n", input)
	fmt.Printf("Output:\n%s\n", string(out))

	subproc.Wait()
}
