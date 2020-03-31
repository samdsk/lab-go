package main

import (
  "fmt"
  "bufio"
  "os"
  _"regexp"
  "strings"
  "strconv"
  _"errors"
  _"sort"
)

type CSV struct{
  nome string
  cognome string
  sesso string
  eta int
}


func main() {
  s := "dati.csv"

  c, err := readCSV(s)

  if err != nil{
    fmt.Println(err)
  }

  fmt.Println(c)

}

func readCSV(nomeFile string) (c []CSV, err error){
  f, err := os.Open(nomeFile)

  defer f.Close()

  if err != nil{
    return nil, err
  }

  c = []CSV{}

  scanner := bufio.NewScanner(f)

  for scanner.Scan(){
    line := scanner.Text()
    v, err := parseCSV(line)
    if err != nil {
      return nil, err
    }
    c = append(c,v)

  }

  //sort.SliceStable(c, func(i,j int)bool{return c[i].eta<c[j].eta})
  return
}


func parseCSV(line string)(c CSV, err error){

  s := strings.Split(line,",")

  /*
  if len(s) != 4{
    err = errors.New("Errore: Riga.")
    return
  }
  */
  /*
  regNome := regexp.MustCompile("[a-Z][a-Z' -]+")
  regSesso := regexp.MustCompile("[mMfF]")
  regEta := regexp.MustCompile("[0-9]+")

  nome := regNome.FindAllString(s[0],-1)
  cognome := regNome.FindAllString(s[1],-1)
  sesso := regSesso.FindAllString(s[2],-1)
  eta, err := strconv.Atoi(regEta.FindAllString(s[3],-1))

  */

  // field := regexp.MustCompile("[A-za-z0-9]+")
  //
  // a := field.FindAllString(line, -1)

  if err != nil{
    return
  }

  nome := s[0]
  cognome := s[1]
  sesso := s[2]
  eta := strings.Trim(s[3]," ")

  c.nome = string(nome)
  c.cognome = string(cognome)
  c.sesso = string(sesso)
  c.eta,_ = strconv.Atoi(eta)


  return
}
