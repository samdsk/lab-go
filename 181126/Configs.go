/* scrivere un programma Go (il file si deve chiamare Config.go) che prenda da stdin un elenco di righe <chiave,valore> nella forma:

chiave =        valore
chiave1 = valore1
chiave2 = valore2
...
chiave2 = nuovoValore2
chiaveN = valoreN

"chiave con spazi" = valoreDiChiaveConSpazi

dove i valori sono numeri decimali (float64) e le chiavi sono stringhe di lunghezza arbitraria

Il programma deve costruire una map e riempirla con i valori trovati in stdin e alla fine stampare tutta la tabella dei valori in forma TABBED, cioè:


chiave		valore
chiave1		valore1
chiave2		nuovoValore2
...
chiaveN		valoreN


*/

package main

import ("fmt")

func main () {
  fmt.Println("Inserire i  dati nel formato Chiave = Valore. Il programma termina facendo seguire 'stop' a un valore (es. chiave = valore stop).")
  m:= make(map[string]float64)
  var s,stop string
  var f float64
    for stop!="stop" {
      fmt.Scanf("%s = %f %s",&s,&f,&stop)
      m[s]=f
    }
    for chiave,valore := range m{
  	fmt.Printf("%s\t%f\n",chiave,valore)
    }
}



package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	var mappa map[string]float64
	mappa=make(map[string]float64)
	scanner:=bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line:=scanner.Text()
		linesplit:=strings.Split(line, " = ")
		number,_:=strconv.ParseFloat(linesplit[1],64)
		mappa[linesplit[0]]=number
	}
	for k,v:=range mappa {
		fmt.Printf("%-10s \t %10.3f\n",k,v)
	}
}
