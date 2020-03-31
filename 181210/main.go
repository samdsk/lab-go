// Definiamo un simulatore per giocare a scacchi
package main

import "fmt"

var s Scacchiera //variabile globale: inizializzata da init

//crea, inizializza e visualizza una scacchiera
func init() {
	s = make(Scacchiera) //creiamo una scacchiera
	primariga := []Tipo{TORRE, CAVALLO, ALFIERE, REGINA, RE, ALFIERE, CAVALLO, TORRE}
	//disponiamo sulle righe #2 e #7 i pedoni bianchi e quelli neri
	//sulle righe #1 e #8 la prima riga bianca e quella nera
	for col := 'a'; col <= 'h'; col++ { //riga 1
		put(s, Pezzo{primariga[col-'a'], BIANCO}, 1, col)
	}
	for col := 'a'; col <= 'h'; col++ { //riga 2
		put(s, Pezzo{PEDONE, BIANCO}, 2, col)
	}
	for col := 'a'; col <= 'h'; col++ { //riga 8
		put(s, Pezzo{primariga[col-'a'], NERO}, 8, col)
	}
	for col := 'a'; col <= 'h'; col++ { //riga 7
		put(s, Pezzo{PEDONE, NERO}, 7, col)
	}
	fmt.Println()
	printChess(s)
}

func main() {
	// per ora esegue un semplice debug delle funzioni base per giocare a scacchi
	r1, c1, r2, c2 := readMove()
	debugMove(r1, c1, r2, c2)
	r1, c1, r2, c2 = readMove()
	debugMove(r1, c1, r2, c2)
	r1, c1, r2, c2 = readMove()
	debugMove(r1, c1, r2, c2)
}

//esegue una mossa, visualizza la tastiera, e mostra il pezzo mosso e quello (eventualmente) rimosso
func debugMove(r1 byte, c1 rune, r2 byte, c2 rune) {
	if ps, pd, ok := Move(s, r1, c1, r2, c2); ok {
		fmt.Printf("muove (%d, %c) in (%d, %c)\n", r1, c1, r2, c2)
		fmt.Printf("mosso:%s rimosso:%s\n", StringPezzo(ps), StringPezzo(pd))
	} else {
		fmt.Println("mossa non valida")
	}
	printChess(s)
}

// visualizza una scacchiera
func printChess(chess Scacchiera) {
	fmt.Printf("%4s", "\u2502")
	for c := 'a'; c < 'a'+8; c++ { // riga 0 con le lettere per denotare le col.
		fmt.Printf("%3c  ", c)
	}
	hline := hline(5) // linea orizzontale
	fmt.Println(hline)
	for r := 8; r != 0; r-- {
		fmt.Printf("%-3d\u2502", r)
		for c := 'a'; c < 'a'+8; c++ {
			t := StringPezzo(chess[Casella{byte(r), c}])
			fmt.Printf("%2s  \u2502", t)
		}
		fmt.Println(hline)
	}
}

// ritorna una linea orizzontale in cui ogni elemento ha dimensione width
func hline(width int) string {
	r := "\n"
	for i := 0; i < 9*width; i++ {
		r += "\u2500"
	}
	return r
}

// legge una mossa nella forma
// r1,c1 r2,c2
// e restituisce òe coordinate di partenza e di arrivo
// bisogna usare gli spazi e la virgola come nell'esempio; digitare newline alla fine
func readMove() (r1 byte, col1 rune, r2 byte, col2 rune) {
	fmt.Print("muovi (r1,c1 r2,c2): ")
	fmt.Scanf("%d,%c %d,%c\n", &r1, &col1, &r2, &col2)
	return
}














