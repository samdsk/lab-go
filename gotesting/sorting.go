package main

import (
  "fmt"
  "sort"
  "io/ioutil"
  "encoding/json"
  "os"
)

func main() {
  m := make(map[string][]int)

  m["a"] = []int{1,2,3,4}
  m["b"] = []int{5,6,7,8}
  m["c"] = []int{8,5,3,7}
  m["d"] = []int{1,8,3,5}
  m["e"] = []int{9,2,6,4}

  j,_ := json.Marshal(m)

  ioutil.WriteFile("prova.json",j,0644)

  s,_ := ioutil.ReadFile("prova.json")

  m1 := make(map[string][]int)

  err := json.Unmarshal(s,&m1)
  if err != nil {
		fmt.Fprintf(os.Stderr, "BRUTTO ERRORE: %v\n", err)
		return
	}

  fmt.Println(m1)

  for _,v := range m1{
    sort.SliceStable(v,func(i,j int)bool{ return v[i]<v[j]})
    fmt.Println(v)
  }


}
