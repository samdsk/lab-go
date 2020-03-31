package main

import "fmt"

//calcolo ricorsivo (efficiente) del n-esimo numero di Fibonacci
//non si considerano possibili overflow
func main() {
	var n int
	fmt.Print("n ? ")
	fmt.Scan(&n)
	fmt.Println(fibo(n))
}

var f = []int{0, 1}

// memorizza nella slice f i valori via via calcolati, in modo da minimizzare le chiamate ricorsive
func fibo(n int) int {
	if l := len(f); n > l {
		return fibo(n-1) + fibo(n-2)
	} else if n == l {
		f = append(f, f[n-1]+f[n-2])
	}
	return f[n]
}
