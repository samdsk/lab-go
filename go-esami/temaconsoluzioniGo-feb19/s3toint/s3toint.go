package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	s3 := os.Args[1]
	val := 0
	const base = 3

	j := 0
	for i := len(s3) - 1; i >= 0; i-- {
		if !isS3(s3[i]) {
			fmt.Println("input non valido")
			return
		}
		val += valS3(s3[i]) * int(math.Pow(base, float64(j)))
		j++
	}
	fmt.Println(val)
}

func isS3(c byte) bool {
	return c == 'z' || c == 'u' || c == 'd'
}

func valS3(c byte) (n int) {
	switch c {
	case 'z':
		n = 0
	case 'u':
		n = 1
	case 'd':
		n = 2
	}
	return
}
