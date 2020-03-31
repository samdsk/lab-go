package main
import(
	"fmt"
	"os"
	"bufio"
)

func main(){
	if len(os.Args)<2{
		fmt.Println("Mancano argomenti")
		return
	}
	filename := os.Args[1]

	f,err:=os.Open(filename)

	if err != nil{
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	vocali := make(map[rune]int)

	for scanner.Scan(){
		line := scanner.Text()
		ContaVocali(line,vocali)
	}

	for k,v:= range vocali{
		fmt.Println(k,":",v)
	}

}

func ContaVocali(s string,vocali map[rune]int){
	for _,k:= range s{
		switch k{
		case 'a':
			vocali[k]++
		case 'A':
			vocali[k]++
		case 'e':
			vocali[k]++
		case 'E':
			vocali[k]++
		case 'i':
			vocali[k]++
		case 'I':
			vocali[k]++
		case 'o':
			vocali[k]++
		case 'O':
			vocali[k]++
		case 'u':
			vocali[k]++
		case 'U':
			vocali[k]++
		}
	}
}
