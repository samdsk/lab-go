/**
Esercizio: calcolatrice RPN (il file deve chiamarsi RPN.go)

https://it.wikipedia.org/wiki/Notazione_polacca_inversa

per vederne una in funzione: http://www.alcula.com/calculators/rpn/

versione semplice:
- 2 accumulatori
- 4 operazioni (+ - * /)
- valori float

versione full: array di accumulatori


Implementazione LOOP:

- legge input utente
 - se numero accumula
 - se operatore collassa

- condizione di uscita (inserimento operatore "q")
- condizioni di errore


A1[]
A2[]

sequenza 3 4 5 6 +

risultato (per iterazioni):

A1[3]
A2[]

A1[4]
A2[3]

A1[5]
A2[4]

A1[6]
A2[5]

A1[11]
A2[]


IMPORTANTE: quando finito (a qualunque livello decidiate di mettervi, se cioè con 2 accum o array) caricate su upload.di.unimi.it che poi facciamo girare i test

*/

/*
package main
import(
    "fmt"
)

// lo "stack" (a due posizioni)
var a1, a2 float32

func main(){
    printStack()
    push(12)
    printStack()
    push(7)
    printStack()
    push(13.5)
    printStack()
    fmt.Println(pop())
    printStack()
}

func pop() float32 {
    r := a1
    a1, a2 = a2, 0
    return r
}

func push(value float32) {
    a1, a2 = value, a1
}

func printStack() {
    fmt.Println(a1, a2)
}
*/

package main
import "fmt"
import "strconv"

func main() {
	var input string
	var accumulatore [4]float64

	for {

		fmt.Scanf("%s", &input)
		in, err := strconv.ParseFloat(input, 64)
 		
		switch {

			case err==nil:

				accumulatore=accumula(in, accumulatore)
				fmt.Println(accumulatore)

			case input=="+":

				a := somma(accumulatore)
				accumulatore=collassa(a, accumulatore)
				fmt.Println(accumulatore)

			case input=="-":

				a := sottrai(accumulatore)
				accumulatore=collassa(a, accumulatore)
				fmt.Println(accumulatore)

			case input=="*":

				a := moltiplica(accumulatore)
				accumulatore=collassa(a, accumulatore)
				fmt.Println(accumulatore)

			case input=="/":

				a := dividi(accumulatore)
				accumulatore=collassa(a, accumulatore)
				fmt.Println(accumulatore)

			case input=="q":
				
				goto end
			
			default:

				fmt.Println("Errore. Il programma verrà terminato.")
				goto end

		}

	}

	end:
}

/*
func isNumber(s string) bool {

	for _, c := range s {

		if !(c=='0' || c=='1' || c=='2' || c=='3' || c=='4' || c=='5' || c=='6' || c=='7' || c=='8' || c=='9') {
			return false
		}
	}

	return true
}
*/

func accumula(a float64, vet [4]float64) [4]float64 {
	var vet1 [4]float64
	
	for i:=len(vet)-1; i>0; i-- {

		vet[i]=vet[i-1]

	}

	vet[0]=a

	vet1=vet
	return vet1
}

func collassa(a float64, vet [4]float64) [4]float64 {
	var vet1 [4]float64

	vet[0]=a
	
	for i:=1; i<len(vet)-1; i++ {

		vet[i]=vet[i+1]

	}

	vet1=vet
	return vet1
}


func somma(vet [4]float64) float64 {
	var r float64

	r=vet[0]+vet[1]

	return r
}


func sottrai(vet [4]float64) float64 {
	var r float64

	r=vet[0]-vet[1]

	return r
}


func moltiplica(vet [4]float64) float64 {
	var r float64

	r=vet[0]*vet[1]

	return r
}

func dividi(vet [4]float64) float64 {
	var r float64

	r=vet[0]/vet[1]

	return r
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func MoveSlice(x []string, y []string, i int) {
	for j := i + 1; j < len(x); j++ {
		x[j] = y[j]
	}
	for k := i - 3; k >= 0; k-- {
		x[k+2] = y[k]
	}
	x = y
} //funzione per riorganizzare le slice dopo un'operazione

func Rpn(n []string) string {
	var m []string
	m = n        //in modo che m non sia vuota
	var last int //l'ultimo numero nello slice m sarà quello dato come soluzione
	var a, b float64

	for _, runa := range n[len(n)-1] {
		if runa == '.' || runa == '0' || runa == '1' || runa == '2' || runa == '3' || runa == '4' || runa == '5' || runa == '6' || runa == '7' || runa == '8' || runa == '9' {
			fmt.Println("Errore: non puoi avere un numero alla fine dell'espressione")
			os.Exit(0)
		}
	} //debug per numero dopo operatori

	for i := 0; i < len(n); i++ {

		if n[0] == "+" || n[0] == "-" || n[0] == "*" || n[0] == "/" {
			fmt.Println("Errore: non puoi avere un operatore all'inizio")
			os.Exit(0)
		} //debug per operatore all'inizio

		for _, runa := range n[i] { //controlla la validità dell'input runa per runa in ciascuna stringa dello slice contenente l'espressione
			if runa == '.' || runa == '0' || runa == '1' || runa == '2' || runa == '3' || runa == '4' || runa == '5' || runa == '6' || runa == '7' || runa == '8' || runa == '9' || n[i] == "+" || n[i] == "-" || n[i] == "*" || n[i] == "/" {

				if len(n) <= 1 {
					break
				}

				switch n[i] { //operazioni, eseguite convertendo avanti e indietro da string a float
				case "+":
					a, _ = strconv.ParseFloat(n[i-2], 64)
					b, _ = strconv.ParseFloat(n[i-1], 64)
					m[i] = strconv.FormatFloat((a + b), 'f', -1, 64)
					MoveSlice(n, m, i)
					i = 0
				case "-":
					a, _ = strconv.ParseFloat(n[i-2], 64)
					b, _ = strconv.ParseFloat(n[i-1], 64)
					m[i] = strconv.FormatFloat((a - b), 'f', -1, 64)
					MoveSlice(n, m, i)
					i = 0
				case "*":
					a, _ = strconv.ParseFloat(n[i-2], 64)
					b, _ = strconv.ParseFloat(n[i-1], 64)
					m[i] = strconv.FormatFloat((a * b), 'f', -1, 64)
					MoveSlice(n, m, i)
					i = 0
				case "/":
					a, _ = strconv.ParseFloat(n[i-2], 64)
					b, _ = strconv.ParseFloat(n[i-1], 64)
					m[i] = strconv.FormatFloat((a / b), 'f', -1, 64)
					MoveSlice(n, m, i)
					i = 0
				}
				last = i
			} else {
				fmt.Println("Errore: non hai inserito un input valido")
				os.Exit(0)
			} // debug dell'input, gestito insieme al for di riga 40
			if n[last] == "." {
				fmt.Println("Errore")
				os.Exit(0)
			} //debug per non inserire il punto (decimali) dopo gli operatori
		}
	}
	return n[last] //restituisce la soluzione (v. riga 23)
}

func main() {
	var term string
	var operation []string
	fmt.Println("Inserire numeri ed operatori separati da uno spazio. Il programma termina con =.")
	fmt.Scan(&term)
	for term != "=" {
		operation = append(operation, term)
		fmt.Scan(&term)
	}
	fmt.Println(Rpn(operation))
}
