 package esami

import (
  "fmt"
  "bufio"
  "errors"
  "os"
  "regexp"
  "strconv"
)

func WriteFile(m Nomi2Voti, fileName string) error{
  f,err := os.Create(fileName)
  defer f.Close()

  if err != nil{
    return err
  }

  for nome, voti := range m{
    fmt.Fprintf(f, "%s",nome)
    for _,v := range voti{
      fmt.Fprintf(f, "\t%d",v)
    }
    fmt.Fprintf(f,"\n")
  }
  return nil
}


func ReadFile(fileName string)(Nomi2Voti,error){
  f,err := os.Open(fileName)
  defer f.Close()

  if err != nil{
    return nil,err
  }

  res := make(Nomi2Voti)

  scanner := bufio.NewScanner(f)

  for scanner.Scan(){
    line := scanner.Text()
    nome, voti, err := ParseLine(line)
    if err != nil{
      return nil, err
    }
    res[nome] = voti
  }
  return res,nil
}

func ParseLine(line string)(nome string,voti Voti,err error){
  regexpNome := regexp.MustCompile("[a-zA-Z][a-zA-Z' -]+")
  regexpVoti := regexp.MustCompile("[0-9]+")
  s := regexpNome.FindAllString(line,-1)

  if len(s) != 1{
    return "",nil,errors.New("Sulla riga " + line + " ci sono zero o più di un nome!")
  }

  nome = s[0]

  v := regexpVoti.FindAllString(line,-1)

  for _,x := range v{
    n, err := strconv.Atoi(x)
    if err !=nil{
      return "",nil,err
    }
    voti = append(voti, n)
  }

  return
}
