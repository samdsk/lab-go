package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DrawPoint(c byte, offset int) (point string) {
	s := strings.Repeat(string(' '), offset)
	point = s + string(c)
	return
}

func DrawSegment(c byte, offset, length int) (segment string) {
	pre := strings.Repeat(string(' '), offset)
	line := strings.Repeat(string(c), length)
	segment = pre + line
	return
}

func main() {
	l, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[2])
	if n <= 0 || l <= 0 {
		return
	}
	c := os.Args[3][0]
	offset := 0
	for i := 1; i <= n; i++ {
		fmt.Println(DrawSegment(c, offset, l))
		offset += l - 1
		for j := 1; j < l; j++ {
			fmt.Println(DrawPoint(c, offset))
		}
	}
}
