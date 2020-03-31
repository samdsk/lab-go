/*
Scrivere un programma go (file Anagrammi.go) che legga due stringhe e restituisca in output un valore di verità
che indica se le due stringhe siano una l'anagramma dell'altra.

Potete provare a farlo in due modi:
- usare le librerie di sistema e lo slicing  (Hint: strings e sort...)
- usare le Map per contare le occorrenze dei caratteri
*/

	/*
	c,1
	i,1
	a,1
	o,1
	
	if val, ok := mymap["key"]; !ok {
    	//do something here
	}
	*/

/*
package main
import "fmt"
func main()  {

	var s1,s2 string
	//var a bool  // di troppo?
	var x,y map[rune]int

	x=make(map[rune]int)
	y=make(map[rune]int)	

	fmt.Print("Inserire la prima stringa: ")
	fmt.Scan(&s1)
	fmt.Print("Inserire la seconda stringa: ")
	fmt.Scan(&s2)

	//a=anagramma(s1,s2,x,y)

	// DONE "ottimizzare"
	if len(s1) == len(s2) && anagramma(s1,s2,x,y) {
		fmt.Println("Le due stringhe sono una l'anagramma dell'altra.")
	}else{
		fmt.Println("Le due stringhe non sono una l'anagramma dell'altra.")
	}
}

func anagramma(s1 string, s2 string, x map[rune]int, y map[rune]int) bool {
	var conta int

	// if len(s1) != len(s2) { return false; }
	for _,r:=range s1{
		x[r]=x[r]+1 // se non c'era nulla cosa ritorna x[r]?  valore di default del tipo mappato
	}

	for _,r:=range s2{
		y[r]=y[r]+1
	}


	for runa1,valore1:=range x{
		//if valore1, ok := y[runa1]; !ok { return false; }

		for runa2,valore2:=range y{ // usare mappa
			if runa1==runa2 && valore1==valore2{
				conta++
			}
		}

	}
	
	return conta==len(x)

	//return true
}
*/




/*
import (
	"fmt"
	"os"
)

func main() {
	var equals bool = true
	var str1, str2 string

	str1, str2 = os.Args[1], os.Args[2]

	occurrencesS1 := map[rune]int{}
	occurrencesS2 := map[rune]int{}

	if equals = len(str1) == len(str2); equals {
		for index, el := range str1 {
			occurrencesS1[el]++
			occurrencesS2[rune(str2[index])]++
		}
		for key, el := range occurrencesS1 {
			if equals = equals && el == occurrencesS2[key]; !equals {
				break
			}
		}
	}

	fmt.Print(str1, " == ", str2, " -> ", equals, "\n")
}*/


/*
package main

import ("fmt")


func main (){

	var a, b string

	var Car map [rune] int32
	Car = make ( map [rune] int32 )
	
	fmt.Scan(&a, &b)

	if len(a)!=len(b){
		fmt.Printf("Falso")
		return
	}

	for _,C:=range(a) {
		Car[C]++
		//fmt.Printf("%c - %d\n", C, Car[C])
	}

	//fmt.Printf("//\n")

	for _,C:=range(b) {
		Car[C]--
		//fmt.Printf("%c - %d\n", C, Car[C])
		if Car[C]<0 {
			fmt.Printf("Falso")
			return
		}
	}

	//fmt.Printf("//\n")

	for C,_:=range(Car) {
		//fmt.Printf("%c - %d\n", C, Car[C])
		if Car[C]!=0 {
			fmt.Printf("Falso")
			return
		}
	}

	fmt.Printf("Vero")

}
*/


package main

import(
    "fmt"
    "strings"
    "sort"
)

func main(){
    var s1, s2 string
    fmt.Scan(&s1, &s2)
    fmt.Println(sorted(s1) == sorted(s2))
}

func sorted(s string) string {
    slc := strings.Split(s, "")
    sort.Strings(slc)
    sortedString := strings.Join(slc, "")
    fmt.Println("sorted ->", sortedString)
    return sortedString
}


/*
package main
import "fmt"
func main() {
	var s1, s2 string
	var mappaParola1, mappaParola2 map[rune]int
	mappaParola1 = make(map[rune]int)
	mappaParola2 = make(map[rune]int)
	fmt.Scan(&s1, &s2)
	parola1 := []rune(s1)
	parola2 := []rune(s2)
	if len(parola1) != len(parola2) {
		fmt.Println("Non sono anagrammi!")
		return
	}
	for _, elem := range parola1 {
		mappaParola1[elem]++
	}
	for _, elem := range parola2 {
		mappaParola2[elem]++
	}
	var incontrato int
	for chiave1, valore1 := range mappaParola1 {
		for chiave2, valore2 := range mappaParola2 {
			if chiave1 == chiave2 {
				incontrato++
			}
			if chiave1 == chiave2 && valore1 != valore2 {
				fmt.Println("Non sono anagrammi")
				return
			}
		}
		if incontrato == 0 {
			fmt.Println("Non sono anagrammi!")
			return
		}
	}
	fmt.Println("Sono anagrammi!")
}


/*
package main
import (
	"fmt"
	"os"
	"reflect"
	)
	
func main(){
	
	stringa1:=os.Args[1]
	stringa2:=os.Args[2]
	
	if len(stringa1)==len(stringa2){
		mappas1:=mappa(stringa1)
		mappas2:=mappa(stringa2)
		eq := reflect.DeepEqual(mappas1, mappas2)
		fmt.Println(eq)
	} else {
		fmt.Println("false")
	}
	
	
}

func mappa(s string)map[rune]int{
		var ripRune map[rune]int 
		ripRune = make(map[rune]int)
		for _,r := range s{
			ripRune[r]=ripRune[r]+1
		}
		return ripRune
	}



*/
