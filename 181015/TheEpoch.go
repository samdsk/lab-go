/*

	Problema
	
	scrivere un programma che calcoli (approssimativamente,
	non importano i bisestili) la data corrispondente
	ad un numero di secondi trascorsi da "the epoch" (1 gen 1970)

	es. @2147483647  che giorno è?

*/



package main

import "fmt"

func main () {
	var sectrascorsi,giorno,mese,anno int
	fmt.Println("Quanto tempo è trascorso da The Epoch?" )
	fmt.Scan(&sectrascorsi)
	giornipassati:=sectrascorsi/60/60/24
	anno=1970+giornipassati/365
	mese=1+(giornipassati%365)/30
	giorno=1+(giornipassati%365)%30
	fmt.Println("E' il", giorno, mese, anno)
}

/*
package main
import "fmt"

func main() {
	var secondiPassati, anniPassati, mesiPassati, giorniPassati, resto int
	fmt.Println("Inserisci i secondi passati da 'The Epoch'")
	fmt.Scan(&secondiPassati)
	
	anniPassati=secondiPassati/31536000
	resto=secondiPassati%31536000
	
	mesiPassati=resto/2592000
	resto=resto%2592000

	giorniPassati=resto/86400

	fmt.Println("Data odierna:", 1+giorniPassati, 1+mesiPassati, 1970+anniPassati)
}









// questo è il problema complementare, dalla data al nr. di secondi
package main

import "fmt"

func main () {
  var d,m,y,da,ma,ya,sec int
  fmt.Println("Inserire la data di oggi nel formato D/M/YYYY")
  fmt.Scanf("%d/%d/%d", &d,&m,&y)
  da= d-1
  ma= m-1
  ya= y-1970
  sec=(da*24*60*60)+(ma*30*24*60*60)+(ya*365*24*60*60) //esiste un modo per contare i mesi da 30 e quelli da 31?
  fmt.Println("Sono passati approssimativamente", sec, "secondi da The Epoch.")
}
*/
