/*
Scrivere un programma Go che prenda in ingresso un float64
e determini se esso sia anche un int.
(Esempio: 5.0 è anche un intero, mentre 5.31 no)

(Opzionale: se la risposta è no arrotondarlo
all'intero più vicino, senza usare il package math)
*/

//Sbagliato: qualsiasi numero che abbia 0 come primo decimale viene considerato Int...
package main
import "fmt"
func main() {
	var numero, numero2 float64

	fmt.Print("Inserire il numero.	")
	fmt.Scan(&numero)

	numero2=numero*10
	//Provato: non funziona con 3.01
	if int(numero2)%10==0 {
		fmt.Println("Il numero è sia Int sia Float.")
	} else {
		fmt.Println("Il numero non è sia Int sia Float.")
	}
}

//Versione corretta?
package main
import "fmt"
func main() {
	var numero float64

	fmt.Print("Inserire il numero.	")
	fmt.Scan(&numero)

	if float64(int(numero))==numero {
		fmt.Println("Il numero è sia Int sia Float.")
	} else {
		fmt.Println("Il numero non è sia Int sia Float.")
	}
}


/*
package main.
import "fmt"
func main() {
	var aFloat float64
	var aInt int
	fmt.Print("Inserisci il valore da controllare ")
	fmt.Scan(&aFloat)
	aInt = int(aFloat)
	if aFloat-float64(aInt) == 0 {
		fmt.Println("Il numero è anche un intero")
	} else {
		fmt.Print("il numero non è un intero ")
		if aFloat-float64(aInt) < 0.5 {
			fmt.Println("e si arrotonda a", aInt)
		} else {
			fmt.Println("e si arrotonda a ", aInt+1)
		}
	}
}

package main
import "fmt"
func main() {
	var numero float64
	fmt.Println("Inserisci numero da controllare:")
	fmt.Scan(&numero)
	var intero int = int(numero)
	var floatintero float64 = float64(intero)
	if numero -floatintero != 0 {
		fmt.Println("Questo numero non è un int")
	}
	if numero - floatintero >= 0.5 {
		fmt.Println("Il numero approssimato è:", floatintero+1)
	} else {
		fmt.Println("Il numero è:", floatintero)
	}
}
*/
