package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)


type Studente struct{
	nome,cognome string
	matr, presenze int
}

func main(){
	filename :=	os.Args[1]
	c := os.Args[2][0]


	f,_ := os.Open(filename)
	defer f.Close()

	var studenti []Studente

	scanner := bufio.NewScanner(f)

	for scanner.Scan(){
		line := scanner.Text()

		p := strings.Fields(line)
		m,_:=strconv.Atoi(p[2])
		pre,_ :=strconv.Atoi(p[3])
		persona := Studente{p[0],p[1],m,pre}

		studenti = append(studenti,persona)

	}

	for _,k:=range studenti{
		if k.nome[0] == c{
			incrementaPresenze(&k)
			
		}
		fmt.Println(k.nome,k.cognome,k.matr,k.presenze)
	}
}

func incrementaPresenze(s *Studente){
	s.presenze = s.presenze + 1
}