package main

import (
  "fmt"
  "bufio"
  _"io"
  _"os"
)

func main() {
  s := "ciao mondo"

  r := []rune(s)

  for _,v := range r{
    if v == 'o'{
      fmt.Println("o")
    }

  }

}

func r(scanner *bufio.Scanner){
  if scanner.Scan(){
    riga := scanner.Text()
    r(scanner)
    fmt.Println(riga)
  }

}


func initSlice(s []int){
  for i:=0;i<len(s);i++{
    s[i] = i
  }
}

func initArray(a *[5]int){
  for i:=0;i<len(a);i++{
    a[i] = i
  }
}
