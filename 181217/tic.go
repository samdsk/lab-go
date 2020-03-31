package main
import (
	"fmt"
)

type Orario struct{
	secondi int
	minuti int
	ore int
}

func main(){
	var s,m,o,secondi int
	
	fmt.Scan(&s,&m,&o,&secondi)
	
	err,h := newOrario(s,m,o)
	
	if err == false{
		fmt.Println("parametri non validi.")
		return 
	}
	
	fmt.Println(h.secondi,":",h.minuti,":",h.ore)
	
	for i:=0;i<secondi;i++{
		tic(&h)	
	}
	
	fmt.Println(h.secondi,":",h.minuti,":",h.ore)
}

func newOrario(s,m,o int) (ok bool, h Orario){
	ok = false
	
	if (s>=0 && s<=59) && (m>=0 && m<=59) && (o>=0 && o<=23){
		ok = true
		h.secondi = s
		h.minuti = m
		h.ore = o
	}
	
	return
}

func tic(h *Orario){
	if h.secondi == 59{
		h.secondi = 0
		if h.minuti == 59{
			h.minuti = 0
			if h.ore == 23{
				h.ore = 0
			}else{
				h.ore += 1			
			}
		}else{
			h.minuti += 1 
		} 
	}else{
		h.secondi += 1
	}

}
