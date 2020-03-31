package main
import(
  "fmt"
  "os"
)

func main(){
  if len(os.Args)<2{
    fmt.Println("manca argomento")
    return
  }
  s := os.Args[1]

  fmt.Println(contaCifre(s))
}


func contaCifre(s string) [10]int{
  var n [10]int
  for _,k := range s{
    switch{
    case k=='0':
      n[0]++
    case k=='1':
      n[1]++
    case k=='2':
      n[2]++
    case k=='3':
      n[3]++
    case k=='4':
      n[4]++
    case k=='5':
      n[5]++
    case k=='6':
      n[6]++
    case k=='7':
      n[7]++
    case k=='8':
      n[8]++
    case k=='9':
      n[9]++

    }
  }

  return n
}
