/*
Dati giorno (gg,mm,aa) e ora (hh,mm) precisa di nascita di qualcuno,
calcolare quando ha raggiunto (o raggiungerà) un miliardo di secondi di età
*/


package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	afterOneBillionSeconds := start.Add(time.Second * 1000000000)

	fmt.Printf("start = %v\n", start)
	fmt.Printf("start.Add(time.Second * one billion) = %v\n", afterOneBillionSeconds)
}
