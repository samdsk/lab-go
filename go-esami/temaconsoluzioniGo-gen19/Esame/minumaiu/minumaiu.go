package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	fileName := os.Args[1]
	reader, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Errore apertura file")
		return
	}

	defer reader.Close()
	r_exp := `([[:upper:]][[:lower:]])+` // minuMAIU

	//tiro su una riga alla volta
	bfdr := bufio.NewReader(reader)
	var allMatches []string
	numLine := 0
	for {
		line, err := bfdr.ReadString('\n')
		if err == io.EOF {
			break
		}
		numLine++
		count := 0
		words := strings.Fields(line)
		for _, s := range words {
			newMatches := estraiParole(s, r_exp)
			if len(newMatches) != 0 {
				count++
				allMatches = append(allMatches, newMatches...)
			}
		}
		fmt.Println("riga", numLine, ":", count, "parola/e con il pattern")
	}
	fmt.Println("stringhe trovate :", allMatches)
}

func estraiParole(testo, r_exp string) (parole []string) {
	var wordRegExp = regexp.MustCompile(r_exp)
	parole = wordRegExp.FindAllString(testo, -1)
	return
}
