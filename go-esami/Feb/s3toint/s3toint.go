package main
import(
	"fmt"
	"os"
	"math"
)


func main(){
	n := os.Args[1]
	sum := 0
	j:=0
	for i:=len(n)-1;i>=0;i--{
		v := val(n[i])
		if v == -1{
			fmt.Println("input non valido")
			return
		}
		sum+=v*int(math.Pow(3,float64(j)))
		j++

	}

	fmt.Println(sum)
}

func val(c byte)int{
	switch c{
	case 'z':
		return 0
	case 'u':
		return 1
	case 'd':
		return 2
	default:
		return -1
	}
}
