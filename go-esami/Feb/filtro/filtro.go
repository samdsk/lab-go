package main
import(
	"fmt"
	"strings"
	"strconv"
	"os"
)

func main(){
	l,_ := strconv.Atoi(os.Args[1])
	n,_ := strconv.Atoi(os.Args[2])
	c := []byte(os.Args[3])[0]

	if l<0 || n<0{
		return
	}

	fmt.Println(DrawSegment(c,0,l))
	lun := l
	for i:=1;i<=n;i++{
		for j:=1;j<l;j++{
			fmt.Println(DrawPoint(c,lun+i-2))
		}
		if i<n{
			fmt.Println(DrawSegment(c,lun+i-2,l))
		}
		lun += l-2

	}


}

func DrawPoint(c byte, k int)string{
	s := strings.Repeat(" ",k)+string(c)
	return s
}

func DrawSegment(c byte, k,l int)string{
	s := strings.Repeat(" ",k)+strings.Repeat(string(c),l)
	return s
}
