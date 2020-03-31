/*
Scrivere un programma che legge un intero positivo n (nr d righe) e
disegna una capanna

  /\
 //\\
//  \\

  caricare su upload.di.unimi.it

*/

package main
import "fmt"

func main(){
	var n int
	fmt.Print("Inserisci n: ")
	fmt.Scan(&n)

	n2 := 2*n

	for i:=n;i>0;i--{
		for j:=0;j<=n2;j++{
			if j > n2-i+1{break}

			if j == i || (j == i+1 && i != n){
				fmt.Print("/")
			}else if j == n2-i+1 || j == n2-i{
				fmt.Print("\\")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
