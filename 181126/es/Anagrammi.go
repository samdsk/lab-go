/*
Scrivere un programma go (file Anagrammi.go) che legga due stringhe e restituisca in output un valore di verità
che indica se le due stringhe siano una l'anagramma dell'altra.

Potete provare a farlo in due modi:
- usare le librerie di sistema e lo slicing  (Hint: strings e sort...)
- usare le Map per contare le occorrenze dei caratteri

nativecall
*/

package main
import(
	"fmt"	
)

func main(){
	var s1,s2 string
	fmt.Print("Inserisci la prima stringa:")
	fmt.Scan(&s1)
	fmt.Print("Inserisci la seconda stringa:")
	fmt.Scan(&s2)
	
	b := anagramma(s1,s2)
	
	fmt.Println(b)
	
}

func anagramma(s1,s2 string)bool{
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	
	if len(s1) != len(s2){return false}
	
	for _,c := range s1{
		m1[c]++
	}
	
	for _,c := range s2{
		m2[c]++
	}
	
	b := false
	
	for k,v := range m1{		
		if v == m2[k]{
			b = true
		}else{return false}
	}

	
	return b
	
}
