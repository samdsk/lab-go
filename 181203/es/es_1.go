/* rifare tac legge da stdin e ribalta

ricorsiva
iterativo


analizzare mem, tempistiche, espressività

time - stampa tempo impiegato per vedere tempo impiegato
/usr/bin/time -v per vedere tutte le informazioni da commando esterno
find / > textInput
cat textInput | ./tac
cat textInput | ./tac > output

wc word count -l per righe

find | es_1.go - pipe per dare input al programma


1.04user 1.89system 0:08.24elapsed 35%CPU
*/

package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	var arr []string
	var s string

	scanner:=bufio.NewScanner(os.Stdin)

  fmt.Print("Inserisci delle stringhe:")
  for scanner.Scan(){
		s = scanner.Text()
		arr = append(arr,s)
	}

  for i:=len(arr)-1;i>=0;i--{
    fmt.Println(arr[i])
  }

  fmt.Println("Done")
}
