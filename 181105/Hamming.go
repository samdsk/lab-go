/*
A) Calcolare la distanza di Hamming(*) tra due stringhe di = lunghezza

B) Calcolare la distanza di ...(**) tra due stringhe di = lunghezza

(*)  numero di posizioni in cui le stringhe differiscono

ciao
ciap
1

ciao
cipp
2



(**) Somma della distanza (valori assoluti) dei caratteri nella stessa posizione

ciao
ciap
0001

a
A
97-65

*/



























package main
import "fmt"

func main() {
	var s1, s2 string
	var distanzaDiHamming, n int

/*
	fmt.Println("Inserire le due stinghe (in caso di stringhe di lunghezza diversa, sarà considerata la lunghezza della più corta):")
	fmt.Scan(&s1, &s2)
	
	if len(s1)<len(s2) {
		n=len(s1)
	} else {
		n=len(s2)
	}
*/

fmt.Println("Inserire le due stinghe:")
	fmt.Scan(&s1, &s2)

	for i:=0; i<len(s1); i++ {
		if s1[i]!=s2[i] {
			distanzaDiHamming++
		}
	}

	fmt.Printf("La distanza di Hamming è %d.\n", distanzaDiHamming)
}
