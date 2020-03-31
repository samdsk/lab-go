// sl-lab.it:11778
// sdid307-050-49:11778

/*
Scrivere un programma che legge un intero positivo n e 
stampa la lettera V

*   *
 * * 
  *
  
  
  caricare su upload.di.unimi.it

*/

package main
import(
	"fmt"
)

func main(){
	var n int
	fmt.Print("Inserisci n: ")
	fmt.Scan(&n)

	for i:=1;i<=n;i++{
		for j:=1;j<=n+n;j++{
			if j == i || j == 2*n-i{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
		
	}
	
}
