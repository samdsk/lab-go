/*
Scrivere un programma che legge un intero positivo n e 
stampa un quadrato di asterischi di lato n.
*/

package main
import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	
	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
	
