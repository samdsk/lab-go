package main

import (
	"fmt"
)

func main() {

	for {
		var s string
		var f float64
		fmt.Println("===")
		fmt.Scanf("%s = %f", &s, &f)
		fmt.Printf("s=%s\n", s)
		fmt.Printf("f=%f\n", f)
	}
}
