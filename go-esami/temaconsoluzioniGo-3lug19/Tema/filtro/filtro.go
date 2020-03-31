package main

import (
	"fmt"
)

func main() {
	var prev, next, diff int
	var trend []byte
	fmt.Scan(&next)
	for {
		prev = next
		_, err := fmt.Scan(&next)
		if err != nil {
			break
		}
		diff = next - prev
		switch {
		case diff > 0:
			trend = append(trend, '+')
		case diff == 0:
			trend = append(trend, '=')
		case diff < 0:
			trend = append(trend, '-')
		}
	}
	fmt.Println(string(trend))
}
