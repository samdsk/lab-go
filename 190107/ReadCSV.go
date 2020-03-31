/*

Creare programma che legga da CSV i dati nella forma:

nome, cognome, sesso, età

e riempie una slice di struct dai campi corrispondenti

poi farne un sort per età  (package 'sort')

(esiste encoding/csv, ma provate a realizzarlo senza usare la lib)

domanda: come impostereste dei test per questo programma?

*/

package main
import "os"
import "regexp"
import "bufio"
import "strconv"
import "sort"
import "fmt"

type persona struct{
	nome string
	cognome string
	sesso string
	età int
}


func main() {
	var persone []persona
	//var int i

	file, _ := os.Open("dati.csv") //manca la chiusura del file aperto

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		field := regexp.MustCompile("[A-za-z0-9]+")

		a := field.FindAllString(line, -1)
		b, _ := strconv.Atoi(a[3])
		persone = append(persone, persona{a[0], a[1], a[2], b})

		//i++
	}

	sort.SliceStable(persone, func(i, j int) bool {return persone[i].età < persone[j].età})

	fmt.Println(persone)
}



package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Persona struct {
	nome    string
	cognome string
	sesso   string
	età     int
}

func Aggiorna(p *[]Persona, s string) {
	sliceriga := strings.Split(s, ", ")
	var Tizio Persona
	Tizio.nome = sliceriga[0]
	Tizio.cognome = sliceriga[1]
	Tizio.sesso = sliceriga[2]
	numero, _ := strconv.Atoi(sliceriga[3])
	Tizio.età = numero
	*p = append(*p, Tizio)
}

func Sorter(p *[]Persona) {
	for indice, _ := range *p {
		for indice2, _ := range *p {
			if (*p)[indice].età < (*p)[indice2].età {
				(*p)[indice], (*p)[indice2] = (*p)[indice2], (*p)[indice]
			}
		}

	}
}

func main() {
	var slice []Persona
	var p *[]Persona
	p = &slice
	f, err := os.Open("input.csv")
	if err != nil {
		fmt.Println("Errore nell'apertura del file")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		riga := scanner.Text()
		Aggiorna(p, riga)
	}
	Sorter(p)
	for i:= 0; i < len(*p); i++ {
		fmt.Printf("%-10s %-10s %-10s %-10d\n", (*p)[i].nome, (*p)[i].cognome, (*p)[i].sesso, (*p)[i].età)
	}
}
