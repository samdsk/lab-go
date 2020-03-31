package main
import "fmt" //importing
//import "math"


func main() {
	//fmt.Println("Ciao Mondo!")

	/*
	var num1, num2 float64
	fmt.Println("inserisci due numeri (float)")
	fmt.Scan(&num1, &num2)
	fmt.Println("somma =", num1+num2)
	fmt.Println("differenza =", num1-num2)
	fmt.Println("prodotto =", num1*num2)
	fmt.Println("quoziente =", num1/num2)
	*/
	/*
	var num1, num2 int64
	fmt.Println("Inserisci dividendo e un divisore")
	fmt.Scan(&num1,&num2)
	fmt.Println("Dividendo: ", num1, " Divisore: ", num2)

	fmt.Println("Quoziente: ", num1/num2)
	fmt.Println("Il resto: ", num1%num2)
	*/

	/*
	var celsius, farenheit float64
	fmt.Println("Inserisci la temperatura in °C da convertire in °F")
	fmt.Scan(&celsius)

	farenheit  = (celsius * 9/5) + 32

	fmt.Println("La temperatura inserita °C: ", celsius, " in Farenheit °F: ", farenheit )
	*/

	/*
	var angoloA, angoloB float64
	fmt.Println("Inserisci i due angoli: ")
	fmt.Scan(&angoloA, &angoloB)
	fmt.Println("Il terzo angolo è: ", 180 - angoloA - angoloB)
	*/

	/*
	var base, altezza float64
	fmt.Println("Inserisci base e altezza ")
	fmt.Scan(&base, &altezza)
	fmt.Println("Il perimetro è: ", 2*(base+altezza), " L'area: ", base*altezza)
	*/

	/*
	var a = b+c
	var b = 7
	//var b = a
	var c = 1

	fmt.Println("Test")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	*/

	/*
	var a, b, c, d, pX, pY float64

	fmt.Println("Inserisci le coordinate del primo punto:")
	fmt.Scan(&a, &b)

	fmt.Println("Inserisci le coordinate del secondo punto:")
	fmt.Scan(&c, &d)

	pX = a-c
	pY = b-d


	fmt.Println("La distanza è: ", math.Sqrt((pX*pX)+(pY*pY)))
	*/

	/*
	var carattere int32
	carattere = 65
	fmt.Println(string(carattere))
	carattere++
	fmt.Println(string(carattere))
	*/

	/*
	var a int32
	fmt.Println("Inserisci ammontare del denaro:")
	fmt.Scan(&a)
	fmt.Println("Inserito: ", a)
	fmt.Println("Banconote da 100:", a/100)
	a = a%100
	fmt.Println("Banconote da 50:", a/50)
	a = a%50
	fmt.Println("Banconote da 20:", a/20)
	a = a%20
	fmt.Println("Banconote da 10:", a/10)
	a = a%10
	fmt.Println("Banconote da 5:", a/5)
	a = a%5
	fmt.Println("Monete da 2:", a/2)
	a = a%2
	fmt.Println("Monete da 1:", a/1)
	*/

	var a, day int64
	fmt.Println("Inserisci la data in secondi (epoch):")

	fmt.Scan(&a)
	fmt.Println("Inserito: ", a)

	day = 3600*24

	fmt.Println("Anni: ", a/(day*365), " Anno: ", 1970 + a/(day*365))
	a = a%(day*365)
	fmt.Println("Mesi: ", a/(day*30))
	a = a%(day*30)
	fmt.Println("Giorni: ", a/day)

}
