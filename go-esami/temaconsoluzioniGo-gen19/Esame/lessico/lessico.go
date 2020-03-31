package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var countLines int
	diz := make(map[string][]int)
	sc := bufio.NewScanner(os.Stdin)
	printMenu()
	for sc.Scan() {
		//fmt.Println("scelta", sc.Text())
		if len(sc.Text()) > 0 {
			switch sc.Text()[0] {
			case 'r':
				fmt.Println("riga:", sc.Text())
			case '+':
				parole := strings.Fields(sc.Text()[1:])
				countLines++
				for _, wrd := range parole {
					diz[wrd] = append(diz[wrd], countLines)
				}
			case '?':
				parola := sc.Text()[2:]
				fmt.Println("parola:", parola)
				fmt.Println("righe", diz[parola])
			case 'p':
				fmt.Println(diz)
			default:
				fmt.Println("invalid option")
			}
		}
	}
}

func printMenu() {
	fmt.Println("+ (legge e memorizza)\n",
		"? (numeri di riga in cui è comparsa la parola data)\n",
		//"l (Stampa il numero di righe lette)\n",
		//"n (Stampa il numero di parole distinte lette)\n",
		"p (Stampa le parole)")
	//"f (Termina l’esecuzione)")
	return
}
