/*
Scrivere un programma che legge un intero positivo n (nr d righe) e
disegna una capanna


  /\
 //\\
//  \\

  caricare su upload.di.unimi.it

*/

/*
package main
import "fmt"

func main(){
	var n int
	fmt.Print("Inserisci n: ")
	fmt.Scan(&n)

	for i:=n;i>0;i--{
		for j:=0;j<=2*n;j++{

			if j == i || (j == i+1 && i != n){
				fmt.Print("/")
			}else if j == 2*n-i+1 || j == 2*n-i{
				fmt.Print("\\")
				if j > 2*n-i{break}
			}else{
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}
*/

/*
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Il programma stampa una capanna di altezza n. Inserire il valore n: ")
	fmt.Scan(&n)
	for r:=0;r<n;r++ {
		for c:=0;c<n*2;c++ {
			if (c==n-r && c!=n) || c==n-1-r {
				fmt.Print("/")
			} else if c==n+r || c==n-1+r {
				fmt.Print("\\")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
*/

package main

import "fmt"

func main() {
	var n int
	fmt.Println("inserire lato n capanna")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Print(" ")
		if i == n-1 {
			fmt.Println("/\\")
		}
	}

	for x := n - 1; x > 0; x-- {
		for y := 0; y < n*2; y++ {
			if x == y {
				fmt.Print("//")
			}
			if y == n*2-x-2 {
				fmt.Print("\\\\")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()

	}
}
