package main
import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

func main(){

	m := make(map[string]float64)
	var s string

	scanner := bufio.NewScanner(os.Stdin)
	for s != "q"{

		fmt.Print("Inserisci chiave = valore (q=exit):")
		scanner.Scan()
		s = scanner.Text()
		temp := strings.Split(s,"=")
		//fmt.Println(temp)

		if temp[0] != "q"{
			m[strings.Trim(temp[0]," ")],_ = strconv.ParseFloat(strings.Trim(temp[1]," "),64)
		}
		//fmt.Println(m)


	}

	for k,v := range m{
		fmt.Println(k,"\t",v)
	}

}
