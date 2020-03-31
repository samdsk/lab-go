/*
Scrivere un programma che legge un intero positivo n e 
stampa un triangolo di asterischi di cateto n.
*/

package main
import(
	"fmt"
)

func main(){
	var x int
	fmt.Print("Inserisc n: ")
	fmt.Scan(&x)

	//crescente
	for i:=1;i<=x;i++{
		for j:=1;j<=i;j++{
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
	//decrescente
	for i:=x;i>0;i--{
		for j:=1;j<=i;j++{
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
