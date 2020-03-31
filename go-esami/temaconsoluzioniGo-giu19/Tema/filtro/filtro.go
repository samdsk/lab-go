package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	min, max := n, n
	sum := n
	count := 1

	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		if n < min {
			min = n
		} else if n > max {
			max = n
		}
		sum += n
		count++
	}
	ave := float64(sum) / float64(count)
	fmt.Println("min:", min)
	fmt.Println("max:", max)
	fmt.Println("media:", ave)
}
