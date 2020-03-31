package main

import(
	"fmt"
)

func main(){
    var n int
    var a,b,c,d float64
    arr := make([]int,0)

    for {
        fmt.Scan(&n)
        if n >= 18 && n <=30{
            arr = append(arr,n)
        }
        if n == 0 {break}
        
    }
    
    l := len(arr)
    for i:=0; i<l; i++{
    
    	if arr[i]>=18 && arr[i]<=21{
    		d += 1
    	}
    	
    	if arr[i]>=22 && arr[i]<=24{
    		c += 1
    	}
    	
       	if arr[i]>=25 && arr[i]<=27{
    		b += 1
    	}
    	
    	if arr[i]>=28 && arr[i]<=30{
    		a += 1
    	}
    
    }
	
	fmt.Println("A :",int((a/float64(l))*100),"%")
	fmt.Println("B :",int((b/float64(l))*100),"%")
	fmt.Println("C :",int((c/float64(l))*100),"%")
	fmt.Println("D :",int((d/float64(l))*100),"%")

}
