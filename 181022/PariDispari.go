/*
Scrivere un programma Go che legge un intero n e, a
seconda del valore di n, stampa uno dei messaggi
	n è pari
oppure
	n è dispari
*/

package main
import "fmt"

func main() {
	var numero int

	fmt.Print("Inserire un numero intero:	")
	fmt.Scan(&numero)

	if numero%2==0 {
		fmt.Println("Il numero inserito è pari.")
	} else if numero%2==1 {
		fmt.Println("Il numero inserito è dispari.")
	}
}







