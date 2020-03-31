package main
import(
	"fmt"
	"math"
	"sort"
	"io"
)

func main(){
	var serie []float64
	var f float64
	for{
		_,err:=fmt.Scan(&f)
		if err == io.EOF{
			break
		}
		serie = append(serie,f)
	}
	fmt.Println("moda:",moda(serie))
	fmt.Println("mediana:",mediana(serie))

}

func moda(serie []float64)float64{
	m := make(map[float64]int)
	for _,k:=range serie{
		m[k]++
	}
	var arr []int
	for _,k:=range m{
		arr = append(arr,k)
	}

	sort.Slice(arr, func(i,j int)bool{
		return arr[i]<arr[j]
	})

	max := arr[len(arr)-1]
	moda := 0.
	count := 0

	for k,v:=range m{
		if count == 0 &&  v == max{
			moda = k
			count++
		}

		if count != 0 && v == max{
			moda = math.Min(moda,k)
		}
	}

	return moda
}

func mediana(serie []float64)float64{

	sort.Slice(serie, func(i,j int)bool{
		return serie[i]<serie[j]
	})

	f := 0.
	if len(serie)%2 == 1{
		f = serie[len(serie)/2]
		return f
	}

	if len(serie)%2 == 0{
		f = (serie[len(serie)/2] + serie[len(serie)/2 -1])/2
		return f
	}

	return f
}
