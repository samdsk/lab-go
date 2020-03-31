package main

import "fmt"

func main() {
	var freq [4]int
	var voto, count int
	for {
		fmt.Scan(&voto)
		if voto <= 0 {
			break
		}
		switch {
		case 18 <= voto && voto <= 21:
			freq[0]++
			count++
		case 22 <= voto && voto <= 24:
			freq[1]++
			count++
		case 25 <= voto && voto <= 27:
			freq[2]++
			count++
		case 28 <= voto && voto <= 30:
			freq[3]++
			count++
		}
	}
	c := 'A'
	for i := 3; i >= 0; i-- {
		fmt.Println(string(c), ":", freq[i]*100/count, "%")
		c++
	}
}
