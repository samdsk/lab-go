package main
import(
	"fmt"
	"math"
	"strconv"
	"os"
	"errors"
)

type Fraction struct{
	num,den int
}

func main(){
	if len(os.Args)!=5{
		fmt.Println("numero di parametri non valido")
		return
	}

	n1,err:=strconv.Atoi(os.Args[1])
	d1,err:=strconv.Atoi(os.Args[2])
	n2,err:=strconv.Atoi(os.Args[3])
	d2,err:=strconv.Atoi(os.Args[4])

	f1,err:=newFraction(n1,d1)

	if err != nil{
		fmt.Println(err)
		return
	}

	f2,err:=newFraction(n2,d2)

	if err != nil{
		fmt.Println(err)
		return
	}

	f:=sum(f1,f2)

	fmt.Println(f.num,"/",f.den)
}

func newFraction(n,d int)(Fraction,error){
	var f Fraction
	if d == 0{
		err := errors.New("divisore nullo")
		return f,err
	}

	min := int(math.Min(float64(n),float64(d)))
	for i:=2;i<=min;{
		if n%i == 0 && d%i == 0{
			n = n/i
			d = d/i
		}else{
			i++
		}
	}

	f.num = n
	f.den = d

	return f,nil
}

func sum(f1,f2 Fraction)Fraction{
	d := f1.den*f2.den
	n := f1.num*f2.den + f2.num*f1.den

	f,_ := newFraction(n,d)

	return f
}
