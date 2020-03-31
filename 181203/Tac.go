/*
rifare 'tac' (legge da stdin e "ribalta") in due modi:

- ricorsivo
- iterativo

analizzare mem, tempistiche, espressività, etc.



Esiti performance:



*/

//Metodo iterativo:

package main
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	var righe []string
	fmt.Println("Inserire il valore stop per interrompere la sequenza di inserimento:")
	scanner:=bufio.NewScanner(os.Stdin)
	
	// leggo file e riempio array
	for scanner.Scan(){
		riga:=scanner.Text()
		righe=append(righe,riga)
	}
	
	// stampo in senso inverso
	for j:=1;j<=len(righe);j++{
		fmt.Println(righe[len(righe)-j])
	}
}

/*
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	files := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		tmp := scanner.Text()
		files = append(files, tmp)
	}

	fmt.Println("Iterative way: ")
	Iterative(files)

	fmt.Println("Recursion way: ")
	Recursion(files)
}

func Iterative(files []string) {
	for i := len(files) - 1; i >= 0; i-- {
		fmt.Println(files[i])
	}
}

func Recursion(files []string) {
	length := len(files) - 1

	if length >= 0 {
		fmt.Println(files[length])
		Recursion(files[:length])
	}
}
*/

package main

import (
    "fmt"
    "bufio"
    "os"
)

//var scanner *bufio.Scanner

func main() {
    scanner:=bufio.NewScanner(os.Stdin)
    recScan(scanner)
}

func recScan(scanner *bufio.Scanner) {
    if scanner.Scan() {
        riga := scanner.Text()
        recScan(scanner);
        fmt.Println(riga)
    }
}
