 	// sl-lab.it:11778
// sdid307-050-49:11778

/*
Scrivere un programma che legge un intero positivo n e 
stampa un triangolo di asterischi di cateto n.

caricare su upload.di.unimi.it

*
**
***
****
*****

*/


package main
import "fmt"

func main() {
	var n int

	fmt.Print("Inserire la lunghezza del cateto: ")
	fmt.Scan(&n)
	
	for i:=1; i<=n; i++ {
		for j:=1; j<=i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}










