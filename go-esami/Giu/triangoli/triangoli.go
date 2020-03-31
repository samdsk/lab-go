package main
import(
	"fmt"
	"math"
	"errors"

)

type Punto struct{
	x,y float64
}

type Triangolo struct{
	P1,P2,P3 Punto
}

func main(){

	var A,B,C Punto
	fmt.Scan(&A.x,&A.y)
	fmt.Scan(&B.x,&B.y)
	fmt.Scan(&C.x,&C.y)

	fmt.Println(lunghezzeLati(A,B,C))

	t,err := newTriangolo(A,B,C)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("triangolo",t)
		fmt.Println("tipo:",tipoTriangolo(t))
	}

}

func lunghezzeLati(A,B,C Punto)[3]float64{
	f := [3]float64{dis(A,B),dis(B,C),dis(A,C)}
	return f
}

func dis(A,B Punto)float64{
	x := A.x-B.x
	y := A.y-B.y
	res := math.Sqrt(x*x+y*y)
	return res
}

func newTriangolo(A,B,C Punto)(Triangolo,error){
	f := lunghezzeLati(A,B,C)
	var t Triangolo
	if f[0]>=f[1]+f[2] || f[1]>=f[0]+f[2] || f[2]>=f[1]+f[0]{
		err := errors.New("non è un triangolo")
		return t,err
	}

	t = Triangolo{A,B,C}
	return t,nil
}

func tipoTriangolo(t Triangolo)int{
	f := lunghezzeLati(t.P1,t.P2,t.P3)


	count := 0
	for i:=0;i<3;i++{
		for j:=i;j<3;j++{
			if math.Abs(f[i]-f[i])<1e-6{
				//fmt.Println(count)
				count++
			}
		}
	}

	return count/2
}
