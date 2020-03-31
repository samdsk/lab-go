/*
func sostituisci(s, s1, s2  []byte, n int ) []byte che date tre stringhe (o meglio, slice di bytes) s,s1,s2
restituisce una slice corrispondente a s in cui l'n-esima occorrenza di s1
viene sostituita con s2,
//ma solo se s1 è una "parola" in s,
//cioè è separata dal resto da degli spazi
(si suggerisce l'uso del package bytes).
Ad esempio se s (vista come stringa) fosse
"nel mio anello c'è dell'oro. nel tuo no"
la sostituzione "nel" -> "Nel" darebbe come risultato
//"Nel mio anello c'è dell'oro. Nel tuo no"
"Nel mio aNellolo c'è dell'oro. Nel tuo no"
*/

package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := os.Args[1]
	old := os.Args[2]
	new := os.Args[3]
	n, _ := strconv.Atoi(os.Args[4])

	//s := "nel mio anello c'è dell'oro. nel   tuo no, ma nello mi ha detto che nel suo c'è. ma nel .."
	//old := "nel"
	//new := "Nel"
	//n := 2
	fmt.Println(s)
	replaced := sostituisci([]byte(s), []byte(old), []byte(new), n)
	fmt.Println(string(replaced))
}

func sostituisci(s, vecchia, nuova []byte, n int) (replaced []byte) {
	posRel := 0
	posAss := 0
	for i := 1; i <= n; i++ {
		posRel = bytes.Index(s[posAss:], vecchia)
		if posRel == -1 {
			replaced = s
			return
		}
		posAss += posRel + len(vecchia)

	}
	posAss -= len(vecchia)
	replaced = append(replaced, s[:posAss]...)
	replaced = append(replaced, nuova...)
	replaced = append(replaced, s[posAss+len(vecchia):]...)
	return
}
