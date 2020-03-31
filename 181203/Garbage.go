package main
import (
	"fmt"
	//"os"
	"runtime"
	//"time"
)

func main() {
	runtime.GC()
	
	var riempi [100000]int
	fmt.Println(riempi[1:10])
	//time.Sleep(time.Second)
	main()
}
