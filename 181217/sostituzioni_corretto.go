package main

import (
	"fmt"
	"os"
	"bytes"
	"strconv"
)


func main(){
	var s, old, nuovo string
	var n int

	s = os.Args[1]
	old = os.Args[2]
	nuovo = os.Args[3]
	n, _= strconv.Atoi(os.Args[4])

	replaced := sostituisci([]byte(s), []byte(old), []byte(nuovo), n)
	fmt.Println(string(replaced))

}


func sostituisci(s,old,nuovo []byte, n int)(res []byte){

	pos := bytes.Index(s,old)
	tot := 1
	oldLen := len(old)

	if pos != -1 && n >0 {

		for i:=n;i>=0;i--{

			if n == tot{
				res = append(res, s[:pos]...)
				res = append(res, nuovo...)
				res = append(res, s[pos+oldLen:]...)
				break
			}

			newPos := bytes.Index(s[pos+oldLen:],old)
			if newPos != -1{
				tot++
				pos = pos + newPos + oldLen
			}else{
				return s
			}
		}

	}else{
		return s
	}

	return res
}
