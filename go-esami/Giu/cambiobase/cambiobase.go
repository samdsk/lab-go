package main
import(
	"fmt"
	"os"
	"strconv"
	"errors"
)

func main(){
	if len(os.Args)!=3{
		fmt.Println("numero parametri non valido")
		return
	}

	n,err := strconv.Atoi(os.Args[1])
	if err != nil{
		fmt.Println("parametro non valido")
		return
	}
	b,err := strconv.Atoi(os.Args[2])
	if err != nil{
		fmt.Println("parametro non valido")
		return
	}

	risultato,err := tenToB(n,b)

	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(risultato)
}


func tenToB(n,b int)(string, error){

	s := ""

	if b>16 || b<2{
		e := errors.New("base non valida")
		return s,e
	}

	val := "0123456789ABCDEF"

	for{
		s = string(val[n%b])+s
		n = n/b
		if n == 0{
			break
		}
	}

	return s,nil
}
