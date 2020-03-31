package main
import(
    "fmt"
    "os"
    "strings"
    "bufio"
)

func main(){

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan(){
        line := scanner.Text()
        let := line[0]

        switch let{
        case 'L':
            L()
        case 'T':
            T()
        case 'Z':
            Z()
        default:
            fmt.Println("input non valido")
    }
    }
}

func L(){
    fmt.Println("*")
    fmt.Println("*")
    fmt.Println("*")
    fmt.Println("*")
    fmt.Println("*****")
}

func T(){
    fmt.Println("*****")
    fmt.Println("  *")
    fmt.Println("  *")
    fmt.Println("  *")
    fmt.Println("  *")
}

func Z(){
    fmt.Println("*****")

    for i:=3;i>0;i--{
        s:=strings.Repeat(" ",i)+"*"
        fmt.Println(s)
    }
    fmt.Println("*****")
}