/*
Scrivere un programma che legge un intero positivo n e 
stampa la lettera V

*   *
 * * 
  *
*/

package main
import(
	"fmt"
)

func main(){
	var n int
	fmt.Print("Inserisci n: ")
	fmt.Scan(&n)
	
	//Righe	
	for i:=1;i<=n;i++{
		//Colonne
		for j:=1;j<=n+n;j++{
			
			if j == i || j == 2*n-i{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
		
	}
	
}
