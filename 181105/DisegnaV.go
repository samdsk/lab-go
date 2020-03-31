// sl-lab.it:11778
// sdid307-050-49:11778

/*
Scrivere un programma che legge un intero positivo n (nr di righe) e 
stampa la lettera V

*       *
 *     *
  *   *
   * * 
    *
  
  
  caricare su upload.di.unimi.it

*/

/*
package main
import "fmt"

func main(){
	var n int
	fmt.Print("Inserire il numero di righe della V: ")
	fmt.Scan(&n)
	for i:=0;i<n;i++{
		for j:=0;j<(n*2)-1;j++{
			if i==j || j==(n*2)-2-i{
				fmt.Print("*")
			}else{
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	
		
}
*/


//Due metodi diversi
package main
import "fmt"

func main() {
	var n int

	fmt.Print("Inserire il numero di righe:	")
	fmt.Scan(&n)
	
	for i:=1; i<n; i++ {
		for j:=0; j<i; j++ {
			fmt.Print("-")
		}
		fmt.Print("*")
		for k:=1; k<2*(n-i); k++ {
			fmt.Print("_")
		}
		fmt.Println("*")
	}

	for j:=0; j<n; j++ {
			fmt.Print(".")
	}
	fmt.Println("*")
}


/*
package main

import "fmt"

func main(){
    var n int
    fmt.Print("Inserire il numero di righe della V: ")
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        for j:= 1; j <= 2*n - i - 1; j++ {
            if j == i {
                fmt.Print("*")
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println("*")
    }
}
*/
