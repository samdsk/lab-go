package main

import (
  "fmt"
  "io"
)

func main(){
  var s []string
  var temp string
  for {
    _, err := fmt.Scan(&temp)
    if err == io.EOF{
      break
    }
    s = append(s,temp)

  }
  fmt.Println(frequenze(s))

}

func frequenze(s []string)map[string]int{
  m := make(map[string]int)

  for _,k := range s{
    m[k]++
  }

  return m
}
