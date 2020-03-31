package main
import(
  "fmt"
  "strconv"
  "os"
)

func main(){
  if len(os.Args)<2 {
    fmt.Println("nessun valore in ingresso")
    return
  }

  prev, err := strconv.Atoi(os.Args[1])

  count := 1
  if err != nil{
    fmt.Println("elemento in posizione",count,"non valido")
    return
  }

  if err == nil && len(os.Args)==2{
    fmt.Println("sequenza valida")
    return
  }


  for i:=2;i<len(os.Args);i++{
    next,err := strconv.Atoi(os.Args[i])

    if err != nil{
      fmt.Println("elemento in posizione",count+1,"non valido")
      return
    }

    if prev<0 && next<0 || prev>=0 && next>=0 {
      fmt.Println("elemento in posizione",count+1,"non valido")
      return
    }
    count++
    prev = next
  }

  fmt.Println("sequenza valida")
}
