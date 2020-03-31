package main
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	scanner:=bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		riga:=scanner.Text()
		fmt.Println(riga)
	}
}
