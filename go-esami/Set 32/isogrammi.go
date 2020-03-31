package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan(){
		line := scanner.Text()
		
		if isIsogram(line){
			fmt.Printf("%s: SI\n",line)
		}else{
			fmt.Printf("%s: NO\n",line)
		}
	}
}

func isIsogram(s string) bool{
	m := make(map[rune]int)
	for _,v:=range s{
		if v =='-'{
			continue
		}else if v == ' '{
			continue
		}else{
			m[v]++
			if m[v]>1{
				return false
			}
		}


	}

	return true
}