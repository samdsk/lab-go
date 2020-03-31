package main
import (
	"fmt"
	"strings"
	"os"
	"os/exec"
	"log"
	"math"
	"time"
	)

/*
Creare un programma che mostri una clessidra che
scandisce un countdown, i parametri di tempo
(minuti) vanno letti da stdin.
Dimensione della clessidra è funzione dei minuti per cui è stata programmata.

Inizio:
**********
 ********
  ******
   ****
   *  *
  *    *
 *      *
*________*

In itinere:
*        *
 ********
  ******
   ****
   *  *
  *    *
 *      *
**********

Fine:
*        *
 *      *
  *    *
   *  *
   ****
  ******
 ********
**********

*/
func main() {

	var durata int
	var carattereRiempimento string
	

// chiedere dati a utente
	fmt.Println("Inserire durata e carattere da usare:")
	fmt.Scan(&durata, &carattereRiempimento)

	/*
	fmt.Println(durata,carattereRiempimento)
	fmt.Println("---------------------------")
	*/

// chiamare funzione disegna clessidra in un ciclo temporizzato,
// cancellando schermo tra un disegno e l'altro

		cancellaSchermo()

	for i:=10;i>0;i-- {
		disegnaClessidra(20,float32(i)/10.0)
		time.Sleep(time.Second)
		cancellaSchermo()
	}
}









// funzione che disegna una singola riga
// accetta lunghezza, stato di piena o meno, carattere da usare, shift
// Dato che avete già visto strings provate a farlo senza for
func disegnaRiga(lunghezza int, spostamento int, stato bool, carattere rune) {
	
/*
	for i:=0; i<spostamento; i++ {
		fmt.Print(" ")
	}
*/

	fmt.Print(strings.Repeat(" ", spostamento))  // sposta il cursore
	fmt.Print(string(carattere)) // stampa bordo clessidra

	if stato {
/*
		for j:=0; j<lunghezza-2; j++ {
			fmt.Print(string(carattere))
*/
	
		fmt.Print(strings.Repeat(string(carattere), lunghezza-2))

	} else {
/*
		for j:=0; j<lunghezza-2; j++ {
			fmt.Print(" ")
*/

		fmt.Print(strings.Repeat(" ", lunghezza-2))
	}
	fmt.Println(string(carattere))  // ultimo bordo (DX)
}


// funzione che disegna tutta la clessidra
// accetta altezza, livello passaggio sabbia (100% sabbia tutta su, 0% tutta giù)
func disegnaClessidra (altezza int, livello float32){

	var larghezza=altezza
	var half=altezza/2
	var carattere='*'

	var rigaRiempimento int
	rigaRiempimento=int(math.Floor(float64(1-livello) * float64(altezza)))
	var shift=0
	for i:=0; i<altezza; i+=1 {
		var fill=i>rigaRiempimento
	 if i<half {
	 	disegnaRiga(larghezza-2*i,shift,fill,carattere)
	 	shift++
	 } else {
	 // TODO mancano i "vuoti" sotto
	 	disegnaRiga(2+larghezza-2*shift,shift-1,fill,carattere)
	 	shift--
	 }
	}

/*
	for i:=	half; i>2; i-=2 {
	 disegnaRiga(i,half-i/2,true,'p')
	}
	for i:=2; i<half; i+=2 {
	 disegnaRiga(i,half-i/2,true,'p')
	}
*/
}
	
// time.Sleep(n_seconds * time.Second)
// (o time.Millisecond, time.Nanosecond....)


/* <--- La mia soluzione
//Disegnare la clessidra
func clessidra(altezzafratto2 int, livelloSabbia float32) {
//il livellosabbia da fornire si ottiene con secondipassati/altezza	
//altezzafratto2 è la metà dell'altezza

	for i:=0; i<altezzafratto2; i++ {
		if float32(i)>=livelloSabbia {
			riga(altezzafratto2*2+1-i*2, i, true, '*')
		} else {
			riga(altezzafratto2*2+1-i*2, i, false, '*')
		}
	}

	for j:=0; j<altezzafratto2; j++ {
		if float32(j)<livelloSabbia {
			riga(1+(1+j)*2, altezzafratto2-j-1, true, '*') 
		} else {
			riga(1+(1+j)*2, altezzafratto2-j-1, false, '*')
		} 

	}
*/


func cancellaSchermo() {
	cmd := exec.Command("clear")
	cmd.Stdout=os.Stdout
	log.Printf("Running command and waiting for it to finish...")
	cmd.Run()
	//log.Printf("Command finished with error: %v", err)
}

/*
package main
import (
	"fmt"
	"time"
	"log"
	"os"
	"os/exec"
)

func main(){
	// time.Sleep(n_seconds * time.Second)
	// (o time.Millisecond, time.NanoSecond....)
	var h int
	var n int
	fmt.Print("Quanto tempo deve durare la clessidra? (secondi)")
	fmt.Scan(&n)
	fmt.Print("Quanto deve essere alta la clessidra? ")
	fmt.Scan(&h)
	for i:=0;i<=n;i++{
		clear()
		Disegna(h,i,n)
		time.Sleep(1*time.Second)
	}
}

func clear(){
	cmd := exec.Command("clear")
	cmd.Stdout=os.Stdout
	log.Printf("Running command and waiting for it to finish...")
	cmd.Run()
	//log.Printf("Command finished with error: %v", err)n()  
}

func Disegna(altezza int,passati int, totale int){
	var vuote int
	vuote=((altezza*passati)/2)/totale
	for i:=0;i<altezza/2;i++{
		if i<vuote{
			newLine(i,altezza+4, false)
		}else{
			newLine(i,altezza+4,true)
		}
		fmt.Println()
	}
	for i:=altezza/2;i>=0;i--{
		if i>vuote{
			newLine(i,altezza+4,false)
		}else{
			newLine(i,altezza+4,true)
		}
		fmt.Println()
	}
}

func newLine(spazio int,lunghezza int,pieno bool){
	if pieno{
		for j:=0;j<spazio;j++{
			fmt.Print(" ")
		}
		for j:=0;j<lunghezza-(spazio*2);j++{
			fmt.Print("*")
		}
	}else{
		for j:=0;j<spazio;j++{
			fmt.Print(" ")
		}
		fmt.Print("*")	
		for j:=0;j<lunghezza-(spazio*2)-2;j++{
			fmt.Print(" ")
		}
		fmt.Print("*")
	}
}
*/

