package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)


func main(){
	scanner := bufio.NewScanner(os.Stdin)
	var s string
	
	for scanner.Scan(){
		s = scanner.Text()
		
	}
	
	temp := strings.Split(s," ")
	
	n,_ := strconv.ParseInt(temp[len(temp)-1],10,64)
	nuovo := temp[len(temp)-2]
	old := temp[len(temp)-3]
	
	stringa := strings.Join(temp[:len(temp)-3]," ")

	

	res := sostituisci([]byte(stringa),[]byte(old),[]byte(nuovo),n)
	fmt.Println(n)
	fmt.Println(stringa)
	fmt.Println(string(res))
}


func sostituisci(s,old,nuovo []byte, n int64)(res []byte){

	res = []byte(strings.Replace(string(s),string(old),string(nuovo),int(n)))
	
	return res
}
