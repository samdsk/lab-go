/* rifare tac legge da stdin e ribalta

ricorsiva
iterativo


analizzare mem, tempistiche, espressività

time - stampa tempo impiegato per vedere tempo impiegato


find | es_1.go - pipe per dare input al programma

ctrl + D per End of File


1.27user 1.87system 0:08.40elapsed 37%CPU

*/


package main

import (
	"fmt"
	"os"
	"bufio"
)
/*
func main() {
	var arr []string
	var s string

	scanner:=bufio.NewScanner(os.Stdin)

  for scanner.Scan(){
		s = scanner.Text()
	  arr = append(arr,s)
	}

  fmt.Println("Reversing...")
  tac(arr)

  fmt.Println("Done")
}


func tac(s []string){
  l := len(s)
  if l <=0{
    fmt.Println("Non hai inserito niente")
    return
  }
  fmt.Println(s[l-1])

  if l-1>0{
    tac(s[:l-1])
  }

}
*/
func main(){
  scanner := bufio.NewScanner(os.Stdin)
  r(scanner)
}

func r(scanner *bufio.Scanner){
  if scanner.Scan(){
    riga := scanner.Text()
    r(scanner)
    fmt.Println(riga)
  }

}
