package main

import (
	"fmt"
)

type Appuntamento struct {
	giorno           int
	h_inizio, h_fine int
}

func main() {
	var gg, h1, h2 int
	agenda := make([]Appuntamento, 0)
	for {
		_, err := fmt.Scan(&gg, &h1, &h2)
		if err != nil {
			break
		}
		a, ok := NewAppuntamento(gg, h1, h2)
		if ok {
			AddAppuntamento(a, &agenda)
		}
	}
	fmt.Println(agenda)
	//fmt.Scan(&gg)
	//fmt.Println(AppuntamentiGiornata(gg, agenda))
}

func IsValidDay(gg int) bool {
	return 1 <= gg && gg <= 366
}

func IsValidTime(h int) bool {
	return 0 <= h && h <= 24
}

func NewAppuntamento(gg, h1, h2 int) (a Appuntamento, ok bool) {
	if IsValidDay(gg) && IsValidTime(h1) && IsValidTime(h2) && h1 <= h2 {
		a = Appuntamento{gg, h1, h2}
		ok = true
	} else {
		ok = false
	}
	return
}

func AddAppuntamento(a Appuntamento, agenda *[]Appuntamento) bool {
	for _, app := range *agenda {
		if CheckSovrapposizioneAppuntamenti(a, app) {
			return false
		}
	}
	*agenda = append(*agenda, a)
	return true
}

func AppuntamentiGiornata(gg int, agenda []Appuntamento) (giornata []Appuntamento) {
	for _, a := range agenda {
		if a.giorno == gg {
			giornata = append(giornata, a)
		}
	}
	return
}

func CheckSovrapposizioneAppuntamenti(a, b Appuntamento) bool {
	if a.giorno != b.giorno {
		return false
	}
	// da qui in poi hanno stesso giorno

	// a inizia in b
	if a.h_inizio >= b.h_inizio && a.h_inizio < b.h_fine {
		return true
	}

	// a finisce in b
	if a.h_fine > b.h_inizio && a.h_fine <= b.h_fine {
		return true
	}

	return false
}
