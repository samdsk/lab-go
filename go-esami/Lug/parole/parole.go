package main
import(
	"fmt"
	"os"
	"bufio"
	"regexp"
)
func main(){
	if len(os.Args)!=2{
		return
	}

	filename := os.Args[1]
	f,err:=os.Open(filename)

	if err != nil{
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var parole []string

	for scanner.Scan(){
		line := scanner.Text()
		parole = append(parole,find(line)...)
	}

	fmt.Println(parole)
	fmt.Println(len(parole))
}

func find(s string)[]string{
	reg := regexp.MustCompile(`[a-zA-Z]+`)
	res := reg.FindAllString(s,-1)
	return res
}
