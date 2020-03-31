package main
import(
	"fmt"
	"io"
)

func main(){

	var n,min,max int
	count := 1
	fmt.Scan(&n)

	min = n
	max = n
	sum := n
	for{
		_,err:=fmt.Scan(&n)

		if err == io.EOF{
			break
		}

		if n<min{
			min = n
		}
		if n>max{
			max = n
		}
		sum+=n
		count++
	}

	fmt.Println("min:",min)
	fmt.Println("max:",max)
	fmt.Println("media:",float64(sum)/float64(count))
}
