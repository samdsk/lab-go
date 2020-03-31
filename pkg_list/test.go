package main

import (
  "fmt"
  "strconv"
  "math/rand"
)

type Person struct{
  Nome string
  Cognome string
  Eta int
  Sesso rune
}

type Student struct{
  Voti []int
  Dati Person
}

func (p Person) String()string{
  eta := strconv.Itoa(p.Eta)
  return "Nome:"+p.Nome+" Cognome:"+p.Cognome+" Eta:"+eta+" Sesso:"+string(p.Sesso)
}

func (s Student) String()string{
  voti := ""

  for _,v := range s.Voti{
    voti += strconv.Itoa(v)+" "
  }

  return s.Dati.Nome+" "+s.Dati.Cognome+" Voti:"+voti
}

func (s *Student) GenVoti(){
  for i:=0;i<15;i++{
    s.Voti = append(s.Voti,rand.Intn(30))
  }
}


func main() {
  p1 := Person{"Mario","Rossi",44,'M'}
  p2 := Person{"Maria","Rossini",46,'F'}
  var s1, s2 Student

  s1.GenVoti()
  s2.GenVoti()

  s1.Dati = p1
  s2.Dati = p2

  fmt.Println(s1,"\n",s2)


}
