package main
import(
	"fmt"
)

type Punto struct{
	x,y float64
}

type Retta struct{
	m,q float64
}


func main(){
	var r1,r2 Retta
	fmt.Scan(&r1.m,&r1.q)
	fmt.Scan(&r2.m,&r2.q)

	var p Punto
	fmt.Scan(&p.x,&p.y)

	if Sotto(p,r1) && Sotto(p,r2){
		fmt.Println(p,"dentro")
	}else{
		fmt.Println(p,"fuori")
	}
}

func Sotto(p Punto, r Retta)bool{
	if p.y <= r.m*p.x+r.q && p.y>=0{
		return true
	}
	return false
}
