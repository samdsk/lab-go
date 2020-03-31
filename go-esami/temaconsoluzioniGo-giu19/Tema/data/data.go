package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var mese = [13]string{"", "gennaio", "febbraio", "marzo",
	"aprile", "maggio", "giugno", "luglio", "agosto",
	"settembre", "ottobre", "novembre", "dicembre"}

func main() {
	validFormat := true
	//formato "2019-06-05 14:11:25"
	s := strings.Split(os.Args[1], " ")
	date := s[0]
	time := s[1]
	if len(date) != 10 || len(time) != 8 {
		validFormat = false
	}

	if validFormat && validTime(time) && validDate(date) {
		//stampa data
		s = strings.Split(date, "-")
		month, _ := strconv.Atoi(s[1])
		fmt.Println(s[2], mese[month], s[0])
	} else {
		fmt.Println("argomento non valido")

	}
}

func validTime(time string) bool {
	t := strings.Split(time, ":")
	if len(t) != 3 {
		return false
	}
	sec, _ := strconv.Atoi(t[2])
	sec_ok := 0 <= sec && sec <= 59
	min, _ := strconv.Atoi(t[1])
	min_ok := 0 <= min && min <= 59
	h, _ := strconv.Atoi(t[0])
	h_ok := 0 <= h && h <= 23
	return sec_ok && min_ok && h_ok
}

func validDate(date string) bool {
	d := strings.Split(date, "-")
	if len(d) != 3 {
		return false
	}
	gg, _ := strconv.Atoi(d[2])
	gg_ok := 1 <= gg && gg <= 31
	mm, _ := strconv.Atoi(d[1])
	mm_ok := 1 <= mm && mm <= 12
	aa, _ := strconv.Atoi(d[0])
	aa_ok := 1900 <= aa && aa <= 2100
	return gg_ok && mm_ok && aa_ok
}
