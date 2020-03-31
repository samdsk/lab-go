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

package main

import (
  "fmt"
  "os"
  "os/exec"
  "strings"
  "time"
  "math"

)

func main() {
  var mins int
  var c string

  fmt.Print("Inserisci numero di secondi: ")
  fmt.Scan(&mins)

   fmt.Print("\nInserisci carattere da usare: ")
   fmt.Scan(&c)

  for i:=mins;i >= 0;i--{
    clessidra(mins,i,c)
    time.Sleep(time.Second)
  }

}

// Disegna linea
// funzione che disegna una singola riga
// accetta lunghezza, stato di piena o meno, carattere da usare, shift
// Dato che avete già visto strings provate a farlo senza for
// l = lunghezza, s = Shift, c = carattere di riempimento, b = stato
func line(l,s int, c string, b bool){
  if b{
    fmt.Print(strings.Repeat(" ",s) + "*" + strings.Repeat(c,int(math.Abs(float64(l-s*2)))) + "*\n")
  }else{
    fmt.Print(strings.Repeat(" ",s) + "*" + strings.Repeat(" ",int(math.Abs(float64(l-s*2)))) + "*\n")
  }
}

// h = hieght, level = livello di sabbia, c = carattere di riempimento
// funzione che disegna tutta la clessidra
// accetta altezza, livello passaggio sabbia (100% sabbia tutta su, 0% tutta giù)
func clessidra(h,level int,c string){
  clear()

  h2 := h*2
  var b bool
  nH := h
  for i:=0;i<=h2;i++{

    if i<h{
      b = (h-i) <= level
      if (h-i) == 0 {b=false}
      line(h2,i,c,b)
    }else{
      b = (h-nH) > level
      line(h2,nH,c,b)
      nH--
    }

  }
  fmt.Println()
}

//Pulisce shell
func clear() {
  //cmd := exec.Command("clear")
	cmd := exec.Command("cmd","/c","cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

}
