package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func TestOccorrenze(t *testing.T) {
	arr := []string{"rr", "iii", "ooo", "iii", "rr"}
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	fmt.Println(arr)
	fmt.Println(frequenze(arr))
}

func TestInput(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./frequenze") // presuppone che sia già stato compilato (dallo script)
	input := "aa bb cc d aa cc aa uu oo pp pippo pluto pippo"
	subproc.Stdin = strings.NewReader(input)
	out, err := subproc.CombinedOutput()

	if err != nil {
		log.Fatalf("Fallito: %s\n", err)
	}
	fmt.Printf("Input:\n%s\n", input)
	fmt.Printf("Output:\n%s\n", string(out))

	subproc.Wait()
}
