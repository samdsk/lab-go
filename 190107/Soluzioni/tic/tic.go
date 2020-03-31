package main

import (
	"fmt"
	"os"
	"strconv"
)

type Orario struct {
	sec int
	min int
	h   int
}

func main() {
	sec, _ := strconv.Atoi(os.Args[1])
	min, _ := strconv.Atoi(os.Args[2])
	h, _ := strconv.Atoi(os.Args[3])

	if ok, orario := newOrario(sec, min, h); ok {
		fmt.Println(StringOrario(orario))
		n, _ := strconv.Atoi(os.Args[4])
		for i := 0; i < n; i++ {
			tic(&orario)
		}
		fmt.Println(StringOrario(orario))

	} else {
		fmt.Println("parametri non validi")
	}
}

func newOrario(sec, min, h int) (ok bool, orario Orario) {
	var ok_sec, ok_min, ok_h bool
	if 0 <= sec && sec < 60 {
		ok_sec = true
	}
	if 0 <= min && min < 60 {
		ok_min = true
	}
	if 0 <= h && h < 24 {
		ok_h = true
	}

	if ok_sec && ok_min && ok_h {
		ok = true
		orario = Orario{
			sec: sec,
			min: min,
			h:   h,
		}
	}
	return
}

func tic(orario *Orario) {
	orario.sec++
	if orario.sec == 60 {
		orario.min++
		orario.sec = 0
		if orario.min == 60 {
			orario.h++
			orario.min = 0
			if orario.h == 24 {
				orario.h = 0
			}
		}
	}
}

func StringOrario(orario Orario) string {
	return fmt.Sprintf("%d:%d:%d", orario.sec, orario.min, orario.h)
}
