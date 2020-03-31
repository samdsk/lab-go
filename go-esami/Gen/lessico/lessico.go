package main
import(
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main(){
  fmt.Print("+ (Legge una riga e memorizza le parole)\n? (Legge una parola e indica le righe che la contengono)\np (Stampa le parole)\n")


  scanner := bufio.NewScanner(os.Stdin)
  dizionario := make(map[string][]int)
  
  count := 1
  for scanner.Scan() {
    line := scanner.Text()


    parole := strings.Fields(line[1:])


    if line[0] == '+'{
      for _,k := range parole{
        dizionario[k] = append(dizionario[k],count)
      }
    }

    if line[0] == 'p'{
      fmt.Println(dizionario)
    }

    if line[0] == '?'{
      fmt.Println("parola:",parole[0])
      fmt.Println("righe:",dizionario[parole[0]])
    }

    count++
  }
}
