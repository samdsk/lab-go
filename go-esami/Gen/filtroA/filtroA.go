package main
import(
  "fmt"
  "os"
  "strconv"
)
func main(){
  if len(os.Args) < 2{
    fmt.Println("nessun valore in ingresso")
    return
  }

  var prev int
  count := 1
  prev,err := strconv.Atoi(os.Args[1])

  if err == nil && len(os.Args)==2{
    fmt.Println("sequenza valida")
    return
  }else if err != nil{
    fmt.Println("elemento in posizione",count,"non valido")
    return
  }



  for i:=2;i<len(os.Args);i++{
    next,err := strconv.Atoi(os.Args[i])

    if err != nil{
      fmt.Println("elemento in posizione",count+1,"non valido")
      return
    }

    if (prev+next)%2 == 0{
      fmt.Println("elemento in posizione",count+1,"non valido")
      return
    }
    count++
    prev = next
  }

  fmt.Println("sequenza valida")
}
