package main

import (
	"fmt"
	"os"
)

func main() {
	r := []rune(os.Args[1])[0]
	st := os.Args[2]
	fmt.Println("occorrenze di", string(r), "in", st)

	fmt.Println(occorrenze(st, r))

}

func occorrenze(st string, r rune) int {
	s := []rune(st)
	pres := 0

	if len(s) == 0 {
		return 0
	}

	if s[0] == r {
		pres = 1
	}

	return pres + occorrenze(string(s[1:]), r)
}
