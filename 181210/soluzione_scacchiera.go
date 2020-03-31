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
		t = string(re - 1 + int(p.tipo)) //-1 perchÃ¨ RE == 1
	}
	return
}

// verifica se le coordinate di una casella sono corrette
func check(r byte, c rune) bool {
	return r > 0 && r < 9 && c >= 'a' && c <= 'h'
}

//le operazioni di base sulla map sono ridefinite per comoditÃ

// ritorna il pezzo che si trova nella casella (r,c) della scacchiera s
// e true oppure false se la casella Ã¨ piena o vuota (in quest'ultimo caso ritorna (Pezzo{},false))
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
// sulla scacchiera s, assumendo che sia il turno di un colore
// ritorna (src, dest, true), dove src e dest sono i pezzi sulle caselle di partenza e di arrivo
// (quest'ultimo sarÃ  uguale al valre di default Pezzo{} se la casella di arrivo Ã¨ vuota)
// ritorna (Pezzo{}, Pezzo{}, false) se la mossa non Ã¨ valida per qualsiasi motivo
func Move(s Scacchiera, turno Colore, r1 byte, c1 rune, r2 byte, c2 rune) (src, dest Pezzo, ok bool) {
	if src, dest, ok = vaildMove(s, turno, r1, c1, r2, c2); ok { // mossa valida
		put(s, src, r2, c2) // dispone il pezzo src sulla casella di destinazione
		remove(s, r1, c1)   // svuota la casella di partenza
	}
	return
}

// verifica se la mossa (r1,b1) in (r2,b2) sulla scacchiera s Ã¨ valida
// controlla che le coordinate delle caselle siano corrette, che la casella iniziale non sia vuota,
// che contenga un pezzo dello stesso colore del turno di gioco,
// che la casella finale sia vuota o contenga un pezzo di colore diverso (attenzione alla regola del pedone!),
// che la mossa sia valida per il pezzo da muovere, in base alle regole,...)
// se la mossa Ã¨ valida ritorna (src, dest, true), dove src e dest sono i pezzi sulle caselle di
// partenza e di arrivo (quest'ultimo sarÃ  uguale a Pezzo{} se la casella di arrivo Ã¨ vuota)
// ritorna (Pezzo{}, Pezzo{}, false) se la mossa non Ã¨ valida
func vaildMove(s Scacchiera, turno Colore, r1 byte, c1 rune, r2 byte, c2 rune) (src, dest Pezzo, ok bool) {
	if ok = check(r1, c1) && check(r2, c2); ok { // le coordinate delle caselle sono ok
		src, ok = get(s, r1, c1)
		if ok = ok && src.colore == turno; ok { //la casella iniziale non Ã¨ vuota ed Ã¨ dello stesso colore del turno di gioco
			var found bool
			dest, found = get(s, r2, c2)
			if ok = !(found && dest.colore == src.colore); ok { // casella di arrivo vuota o contenente pezzo di colore diverso
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
		}
	}
	return
}

// differenza in valore assoluto ("safe") tra valori byte (unsigned!); risultato come int
func absDiffB(b1, b2 byte) int {
	if b1 >= b2 {
		return int(b1 - b2)
	}
	return int(b2 - b1)
}

// differenza in valore assoluto tra rune (definita solo per convenienza); risultato come int
func absDiffR(r1, r2 rune) int {
	if r1 >= r2 {
		return int(r1 - r2)
	}
	return int(r2 - r1)
}

//verifica che un cammino sulla scacchiera s che parte dalla casella (r1,c1) non contiene pezzi
//il cammino viene determinato da un numero di passi (steps) in ciascuno dei quale si effettua uno spostamento
//orizzontale pari a dx ed uno verticale pari a dy
func checkPath(s Scacchiera, r1 byte, c1 rune, dx, dy, steps int) bool {
	for ; steps > 1; steps-- {
		c1 += rune(dx)
		r1 += byte(dy)
		if _, found := get(s, r1, c1); found {
			return false
		}
	}
	return true
}

// ritorna 1 o -1, a seconda che cond sia true o false
func singleStep(cond bool) int {
	if cond {
		return 1
	}
	return -1
}

func checkAlfiere(s Scacchiera, r1 byte, c1 rune, r2 byte, c2 rune) (ok bool) {
	steps := absDiffB(r1, r2)
	if ok = steps == absDiffR(c1, c2); ok { // spostamento in diagonale
		dx, dy := singleStep(c1 < c2), singleStep(r1 < r2) //singolo spostamento orizzontale e verticale
		ok = checkPath(s, r1, c1, dx, dy, steps)
	}
	return
}

func checkTorre(s Scacchiera, r1 byte, c1 rune, r2 byte, c2 rune) (ok bool) {
	dy, dx := absDiffB(r1, r2), absDiffR(c1, c2)
	if ok = dx == 0 || dy == 0; ok {
		steps := dx + dy
		if dy == 0 { //orizzontale
			dx = singleStep(c1 < c2)
		} else { // verticale
			dy = singleStep(r1 < r2)
		}
		ok = checkPath(s, r1, c1, dx, dy, steps)
	}
	return
}

func checkRe(r1 byte, c1 rune, r2 byte, c2 rune) bool {
	return absDiffB(r1, r2) == 1 || absDiffR(c1, c2) == 1
}

func checkCavallo(color Colore, r1 byte, c1 rune, r2 byte, c2 rune) (ok bool) {
	if !prima[color] { //non Ã¨ la prima mossa per questo colore
		dy, dx := absDiffB(r1, r2), absDiffR(c1, c2)
		ok = dx == 1 && dy == 2 || dx == 2 && dy == 1
	}
	return
}

func checkPedone(s Scacchiera, color Colore, piena bool, r1 byte, c1 rune, r2 byte, c2 rune) bool {
	dx := absDiffR(c1, c2)
	return (dx == 1 && piena || dx == 0 && !piena) &&
		(color == BIANCO && r2 == r1+1 || color == NERO && r2 == r1-1)
}
