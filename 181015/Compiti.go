/* Quoziente e resto

Problema Scrivere un programma Go che, dati un dividendo e un divisore
(interi), calcoli il quoziente e il resto.

Annotazioni L’operatore per la divisione (/) tra interi calcola la parte intera
del risultato. L’operatore per il resto della divisione è %.

Cosa succede se il divisore è 0?

e se sono entrambi 0?
*/

package main

import "fmt"

func main() {
	var dividendo, divisore int
	fmt.Println("Inserire dividendo e divisore:")
	fmt.Scan(&dividendo, &divisore)
	
	fmt.Println("Quoziente:", dividendo/divisore)
	fmt.Println("Resto:", dividendo%divisore)
}

//Programma con controllo sul divisore

package main

import "fmt"

func main() {   
        var dividendo, divisore int
        var isZero bool
        fmt.Println(" Inserisci prima il dividendo ")
        fmt.Scan(&dividendo)
        if divisore==0 {
                isZero = true
        }
        fmt.Println(" Inserisci il divisore ")
        fmt.Scan(&divisore)
        for isZero {
                fmt.Println(" Inserisci il divisore ")
                fmt.Scan(&divisore)
                if divisore!=0 {
                        isZero = false
                }
        }
        fmt.Println(" Il quoziente e' : ", dividendo/divisore)
        fmt.Println(" Il resto e' : ", dividendo%divisore)
}
