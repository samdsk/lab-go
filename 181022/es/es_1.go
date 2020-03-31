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
	var x int

	fmt.Print("Inserire un numero:	")
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println("Il numero ", x ," è pari.")
	} else if x%2 == 1 {
		fmt.Println("Il numero ", x ," è dispari.")
	}
}
