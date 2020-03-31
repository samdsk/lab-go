package main

//tipi e funzioni base per implementare una scacchiera

type Colore bool // alias per il colore dei pezzi
type Tipo byte   //alias per il tipo dei pezzi

const (
	BIANCO = false
	NERO   = true
)

//i simboli dei pezzi sono disposti nello stesso ordine in cui compaiono in tab Unicode ...
const (
	NULL = iota // 0
	RE
	REGINA
	TORRE
	ALFIERE
	CAVALLO
	PEDONE
)

type (
	Pezzo struct {
		tipo   Tipo
		colore Colore
	}

	Casella struct {
		riga byte //riga: 1..8
		col  rune //colonna: 'a'..'h'
	}

	Scacchiera map[Casella]Pezzo
)

// restituisce la rapresentazione testuale (t) di un Pezzo degli scacchi
// se il pezzo non esiste restituisce ""
func StringPezzo(p Pezzo) (t string) {
	re := 0x2654 // "re bianco" in Unicode; i valori dei pezzi sono contigui (prima i bianchi poi i neri)
	if p.tipo >= RE && p.tipo <= PEDONE {
		if p.colore == NERO {
			re += 6
		}
		t = string(re - 1 + int(p.tipo)) //-1 perchè RE == 1
	}
	return
}

// verifica se le coordinate di una casella sono corrette
func check(r byte, c rune) bool {
	return r > 0 && r < 9 && c >= 'a' && c <= 'h'
}

//le operazioni di base sulla map sono ridefinite per comodità

// ritorna il pezzo che si trova nella casella (r,c) della scacchiera s
// e true oppure false se la casella è piena o vuota (in quest'ultimo caso ritorna (Pezzo{},false))
func get(s Scacchiera, r byte, c rune) (p Pezzo, found bool) {
	p, found = s[Casella{r, c}]
	return
}

// dispone il pezzo p sulla casella (r,c) della scacchiera s
// non esegue alcun controllo
func put(s Scacchiera, p Pezzo, r byte, c rune) {
	s[Casella{r, c}] = p
}

// svuota la casella (r,c) della scacchiera s
// non esegue alcun controllo
func remove(s Scacchiera, r byte, c rune) {
	delete(s, Casella{r, c})
}

// se possibile, in base alla attuale disposizione dei pezzi, esegue la mossa (r1,b1) in (r2,b2)
// sulla scacchiera s
// ritorna (src, dest, true), dove src e dest sono i pezzi sulle caselle di partenza e di arrivo
// (quest'ultimo sarà uguale al valre di default Pezzo{} se la casella di arrivo è vuota)
// ritorna (Pezzo{}, Pezzo{}, false) se la mossa non è valida per qualsiasi motivo
func Move(s Scacchiera, r1 byte, c1 rune, r2 byte, c2 rune) (src, dest Pezzo, ok bool) {
	if src, dest, ok = vaildMove(s, r1, c1, r2, c2); ok { // mossa valida
		put(s, src, r2, c2) // dispone il pezzo src sulla casella di destinazione
		remove(s, r1, c1)   // svuota la casella di partenza
	}
	return
}

// verifica se la mossa (r1,b1) in (r2,b2) sulla scacchiera s è valida
// controlla che le coordinate delle caselle siano corrette, che la casella iniziale non sia vuota,
// che la casella finale sia vuota o contenga un pezzo di colore diverso (attenzione alla regola del pedone!),
// che la mossa sia valida per il pezzo da muovere, in base alle regole,...)
// se la mossa è valida ritorna (src, dest, true), dove src e dest sono i pezzi sulle caselle di
// partenza e di arrivo (quest'ultimo sarà uguale a Pezzo{} se la casella di arrivo è vuota)
// ritorna (Pezzo{}, Pezzo{}, false) se la mossa non è valida
func vaildMove(s Scacchiera, r1 byte, c1 rune, r2 byte, c2 rune) (src, dest Pezzo, ok bool) {
	if ok = check(r1, c1) && check(r2, c2); ok { // le coordinate delle caselle sono ok
		if src, ok = get(s, r1, c1); ok { //la casella iniziale non è vuota
			//var found bool
			dest, found = get(s, r2, c2)
			if ok = !(found && dest.colore == src.colore+); ok { // casella di arrivo vuota o contenente pezzo di colore diverso
				switch src.tipo {
				case ALFIERE:
					ok = checkAlfiere(s, r1, c1, r2, c2)
				case TORRE:
					ok = checkTorre(s, r1, c1, r2, c2)
				case REGINA:
					ok = checkAlfiere(s, r1, c1, r2, c2) || checkTorre(s, r1, c1, r2, c2)
				case CAVALLO:
					ok = checkCavallo(src.colore, r1, c1, r2, c2)
				case PEDONE:
					ok = checkPedone(s, src.colore, found, r1, c1, r2, c2)
				case RE:
					ok = checkRe(r1, c1, r2, c2)
				}
			}
			//completare ...
		}
	}
	return // implementazione fittizia -- da completare
}


// aggiungere funzioni...

// SOLUZIONE: https://homes.di.unimi.it/~capra/labprog1819/lezioni/1001/
