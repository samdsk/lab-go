package main

import "fmt"

func init() {
 fmt.Println("ciao1")
}

func main(){
 //init()  // non dovrebbe compilare, infatti restituisce "undefined func"
}



func init() {
 fmt.Println("ciao3")
}


func init() {
 fmt.Println("ciao2")
}
