/*
Scrivere un programma che legge un intero positivo n e 
stampa un quadrato di asterischi di lato n.
*/

package main
import(
	"fmt"
)

func main(){
	var x int
	fmt.Print("Inserisc n: ")
	fmt.Scan(&x)

	for i:=0;i<x;i++{
		for j:=0;j<x;j++{
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

// |wc -l word count
