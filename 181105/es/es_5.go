/*
Dati giorno (gg,mm,aa) e ora (hh,mm) precisa di nascita di qualcuno,
calcolare quando ha raggiunto (o raggiungerà) un miliardo di secondi di età
*/

package main
import (
	"fmt"
	"time"
)

func main(){
	var anno,mese,giorno,ore,min int


	fmt.Print("Inserisci la data e ora precisa della nascità: [anno mese(01-12) giorno ora min] ")
	fmt.Scan(&anno,&mese,&giorno,&ore,&min)

	oggi := time.Now()

	start := time.Date(anno, time.Month(mese), giorno, ore, min, 0, 0, time.UTC)
	billionSeconds := start.Add(time.Second * 1000000000)
	
	if billionSeconds.Unix() < oggi.Unix() {
		fmt.Println("Hai raggiunto un miliardo di secondi: ",billionSeconds)
	}else{
		fmt.Println("Raggiungerai un miliardo di secondi: ",billionSeconds)
	}


}
