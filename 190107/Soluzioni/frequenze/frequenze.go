package main

import (
	"fmt"
	"io"
)

func main() {
	var st string
	var sl []string
	for {
		_, err := fmt.Scan(&st)
		if err == io.EOF {
			break
		}
		sl = append(sl, st)
	}

	freq := frequenze(sl)

	fmt.Println(freq)
}

func frequenze(s []string) map[string]int {
	freq := make(map[string]int)
	for _, st := range s {
		freq[st]++
	}
	return freq
}
