package main
import(
	"fmt"
	"io"
	"math"
)

type Punto struct{
	x,y float64
}

func main(){

	var tappe []Punto
	for{
		var p Punto
		_,e := fmt.Scan(&p.x,&p.y)

		if io.EOF == e{
			break
		}


		tappe = append(tappe,p)
	}

	fmt.Println("tappe:",tappe)
	fmt.Println("percorso",percorso(tappe))

}

func distanza(p1, p2 Punto)float64{
	x := p1.x-p2.x
	y := p1.y-p2.y
	r := math.Sqrt(x*x+y*y)
	return r
}

func percorso(tappe []Punto)float64{
	if len(tappe)<2{
		return 0
	}
	sum := 0.
	for i:=0;i<len(tappe)-1;i++{
		sum+=distanza(tappe[i],tappe[i+1])
	}

	sum += distanza(tappe[len(tappe)-1],tappe[0])

	return sum
}
