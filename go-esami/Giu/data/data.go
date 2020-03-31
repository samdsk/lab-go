package main
import(
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	if len(os.Args)<2{
		fmt.Println("argomento non valido")
		return
	}

	s := strings.Fields(os.Args[1])

	date := strings.Split(s[0],"-")

	if !Date(date) || len(date)!=3 {
		fmt.Println("argomento non valido")
		return
	}

	time := strings.Split(s[1],":")

	if len(time)!=3 || !Time(time) {
		fmt.Println("argomento non valido")
		return
	}
	m,_ := strconv.Atoi(date[1])
	mesi :=  []string{"gennaio","febraio","marzo","aprile","maggio","giugno","luglio","agosto","settembre","ottobre","novembre","dicembre"}

	fmt.Println(date[2],mesi[m-1],date[0])
}

func Date(d []string)bool{
	b := false
	if d[0]>="1900" && d[0]<="2100" && d[1]>="01" && d[1]<="12" && d[2]>="01" && d[2]<="31"{
		b = true
	}
	return b
}

func Time(t []string)bool{
	b := false
	if len(t)!=3{
		return b
	}
	if t[0]>="00" && t[0]<="23" && t[1]>="00" && t[1]<="59" && t[2]>="00" && t[2]<="59"{
		b = true
	}
	return b
}
