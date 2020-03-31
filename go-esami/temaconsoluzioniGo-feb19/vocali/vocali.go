package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Mancano argomenti")
		return
	}

	nomeFile := os.Args[1]
	reader, err := os.Open(nomeFile)

	if err != nil {
		fmt.Println("Errore apertura file")
		return
	}

	defer reader.Close()

	contaVocali := make(map[rune]int)

	// una riga alla volta, inutile avere tutto il file
	// volendo si può leggere una stringa alla volta
	bfdr := bufio.NewReader(reader)
	for {
		line, err := bfdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		ContaVocali(line, contaVocali)
	}

	for k, v := range contaVocali {
		fmt.Println(string(k), ":", v)
	}
}

func ContaVocali(s string, vocali map[rune]int) {
	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vocali[r]++
		}
	}
}
