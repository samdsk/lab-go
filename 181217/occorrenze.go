package main

import (
  "fmt"
  "os"
)

func main(){
  r := []rune(os.Args[1])[0]
  s := os.Args[2]

  fmt.Println("Occorrenze di",string(r),"in",s)
  fmt.Println(occorrenze(r,s))

}

func occorrenze(r rune, st string)(n int){

  s := []rune(st)

  if len(s) <= 0{
    return n
  }else{
    if r == s[0]{
      n = 1
    }
    n += occorrenze(r,st[1:])
  }

  return n
  // si poteva fare anche cosi
  //return n + occorrenze(r,st[1:])
}
