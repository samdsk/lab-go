package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

// TODO per ora test alla cieca (manca ancora implementazione), vedi modifica del testo del tema esame!

func TestSostituzioni1(t *testing.T) {
	s := "   nel mio anello c'è dell'oro. nel   tuo no, ma nello mi ha detto che nel suo c'è. ma nel .."
	replaced := sostituisci([]byte(s), []byte("nel"), []byte("Nello"), 2)
	ok := "   nel mio aNellolo c'è dell'oro. nel   tuo no, ma nello mi ha detto che nel suo c'è. ma nel .."
	if strings.Compare(string(replaced), ok) != 0 {
		t.Fail()
	}
}

func TestSostituzioni2(t *testing.T) {
	s := "nel mio anello c'è dell'oro. nel   tuo no, ma nello mi ha detto che nel suo c'è. ma nel"
	replaced := sostituisci([]byte(s), []byte("nel"), []byte("Nel"), 2)
	ok := "nel mio aNello c'è dell'oro. nel   tuo no, ma nello mi ha detto che nel suo c'è. ma nel"
	if strings.Compare(string(replaced), ok) != 0 {
		t.Fail()
	}
}

func TestMainComeDaEsempio(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./sostituzioni", "nel mio anello c'è dell'oro. nel    tuo no", "nel", "Nel", "2") // presuppone che sia già stato compilato (dallo script)
	out, err := subproc.CombinedOutput()

	if err != nil {
		//log.Fatalf("Fallito: %s\n", err)
		fmt.Printf("Fallito: %s\n", err)
	}
	fmt.Printf("Command line:\n%+q\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	subproc.Wait()
}

func TestMainAltro(t *testing.T) {
	fmt.Println("N.B. nessun controllo sulla validità dell'output, verificare a vista")
	subproc := exec.Command("./sostituzioni", "uno due tre quattro cinque uno due tre quattro cinque", "tre", "TRE", "2") // presuppone che sia già stato compilato (dallo script)
	out, err := subproc.CombinedOutput()

	if err != nil {
		//log.Fatalf("Fallito: %s\n", err)
		fmt.Printf("Fallito: %s\n", err)
	}
	fmt.Printf("Command line:\n%+q\n", subproc.Args)
	fmt.Printf("Output:\n%s\n", string(out))

	subproc.Wait()
}
