package main
import(
	"fmt"
	"io"
)

func main(){

	var next,prev int
	var s string
	fmt.Scan(&prev)

	for{
		_,err := fmt.Scan(&next)
		if err == io.EOF{
			break
		}

		if next>prev{
			s+="+"
		}
		if next==prev{
			s+="="
		}
		if next<prev{
			s+="-"
		}

		prev = next
	}

	fmt.Println(s)
}
