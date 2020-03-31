package main

import (
	"fmt"
	"os"
)

func main() {
	frase := os.Args[1]
	//frase := "1, 2, 3; non c’è fiaba senza re; 1, 2, 3; venite giù da me; 4, 5, 6; siete dei babbei; 7,8,9; io sono già altrove."
	fmt.Println(contaCifre(frase))
}

func contaCifre(s string) (numCifre [10]int) {
	for _, r := range s {
		if '0' <= r && r <= '9' {
			i := r - '0'
			numCifre[i]++
		}
	}
	return
}
