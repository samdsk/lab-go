package main

import(
	"fmt"
	"os"
	"strconv"
)

func main(){
	if len(os.Args) < 2{
		fmt.Println("no input")
		return
	}

	for i:=1;i<=len(os.Args)-1;i++{

		_,err := strconv.Atoi(os.Args[i]) 

		
		if len(os.Args[i]) != 16 || err != nil{
			
			fmt.Println(os.Args[i],"dato non valido")
			
		}else{
			var valore [16]int
			dato := os.Args[i]
			for j:=0;j<16;j++{
				temp,_ := strconv.Atoi(string(dato[j]))

				valore[j]=temp
			}

			if luhn(valore){
				fmt.Println(os.Args[i],"valido")
			}else{
				fmt.Println(os.Args[i],"non valido")
			}
		}
	}

}

func luhn(n [16]int)bool{

	sum := 0
	for i:=14;i>=0;i = i-2{
		temp := n[i]*2

		if temp >= 10{
			temp = temp - 9

			n[i] = temp
		}else{
			n[i] = temp
		}		
	}

	for _,k:=range n{
		sum += k
	}

	if sum%10 == 0{
		return true
	}

	return false
}