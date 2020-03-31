package main
import(
	"fmt"
	"os"
	"regexp"
	"bufio"
)

func main(){
	if len(os.Args)<2{
		return
	}
	filename := os.Args[1]

	f,e := os.Open(filename)
	defer f.Close()
	if e != nil{
		fmt.Println(e)
		return
	}

	scanner := bufio.NewScanner(f)
	riga := 1
	var s []string

	r := "([[:lower:]][[:upper:]])+"

	for scanner.Scan(){
		line := scanner.Text()
		s = append(s,estraiParole(line,r)...)
		fmt.Println("riga",riga,":",len(estraiParole(line,r)),"parola/e con il pattern")
		riga++
	}

	fmt.Println("stringhe trovate :",s)
}

func estraiParole(testo, r_exp string)(parole []string){
	r := regexp.MustCompile(r_exp)
	parole = r.FindAllString(testo,-1)
	return
}
