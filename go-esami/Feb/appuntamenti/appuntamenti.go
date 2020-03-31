package main
import(
	"fmt"
	"io"
)

type Appuntamento struct{
	giorno, oraInizio, oraFinale int
}


func main(){
	var gg,h1,h2 int
	var agenda []Appuntamento

	for{
		_,err:=fmt.Scan(&gg,&h1,&h2)

		if err == io.EOF{
			break
		}

		app,ok := NewAppuntamento(gg,h1,h2)

		if ok{
			AddAppuntamento(app,&agenda)
		}

	}

	fmt.Println(agenda)

}

func NewAppuntamento(gg,h1,h2 int)(Appuntamento,bool){
	var a Appuntamento

	if gg>366 || gg<1 || h1<0 || h1>24 || h2<0 || h2>24 || h1>h2{
		return a,false
	}

	a = Appuntamento{gg,h1,h2}
	return a,true
}

func CheckSovrapposizioneAppuntamenti(app1,app2 Appuntamento)bool{
	if app1.giorno != app2.giorno{
		return false
	}

	if app1.oraInizio>=app2.oraInizio && app1.oraInizio<app2.oraFinale{
		return true
	}

	if app1.oraFinale>app2.oraInizio && app1.oraFinale<=app2.oraFinale{
		return true
	}

	return false
}

func AddAppuntamento(app Appuntamento,agenda *[]Appuntamento)bool{
	for _,k := range *agenda{
		if CheckSovrapposizioneAppuntamenti(app,k){
			return false
		}
	}

	*agenda = append(*agenda,app)
	return true
}

func AppuntamentiGiornata(gg int,agenda []Appuntamento)(app []Appuntamento){
	for _,k:=range agenda{
		if k.giorno == gg{
			app = append(app,k)
		}
	}

	return
}
