package main
import "fmt"

var(
  s Scacchiera
  prima map[Colore]bool = map[Colore]bool{BIANCO: true, NERO: true}
)

func init(){
  s = make(Scacchiera)
  primariga := []Tipo{TORRE, CAVALLO, ALFIERE, REGINA, RE, ALFIERE, CAVALLO, TORRE}

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
}
func main(){
  fmt.Println("semplice simulatore per gli scacchi\nsintassi mossa: r1,c1 r2,c2")
  fmt.Println("r1,r2 compresi tra 1..8; c1,c2 tra 'a'..'h' (non inserire spazi tra rx e cx; separare con ',' - premere invio alla fine)")
  fmt.Println("Ctrl-c per terminare in anticipo")
  fmt.Println("per una corretta visualizzazione impostare il terminale come utf-8 e con lo sfondo bianco")
  printChess(s)

  namecol := map[Colore]string{BIANCO: "Bianco", NERO: "Nero"}
  turno := Colore(BIANCO) //inizia il bianco

  for mangiato := Tipo(NULL); mangiato != RE; {
    fmt.Printf("muove il %s: ", namecol[turno])
    r1, c1, r2, c2 := readMove()

    if src, dest, ok := Move(s, turno, r1,r2,c1, c2); ok {
      fmt.Printf("mosso:%s rimosso:%s\n", NomePezzo(src), NomePezzo(dest))
      mangiato = dest.tipo
      prima[turno] = false
      turno = !turno
    }	else {
  			fmt.Println("mossa non valida")
  	}

    printChess(s)
  }
  fmt.Printf("scacco! vince il %s\n", namecol[!turno])
}


func readMove() (r1 byte, col1 rune, r2 byte, col2 rune) {
	fmt.Scanf("%d%c %d%c\n", &r1, &col1, &r2, &col2)
	return
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
			t := NomePezzo(chess[Casella{byte(r), c}])
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
