package main

import (
    "fmt"
)

func main() {
    for i :=1 ; i < 1000; i++ {
        fmt.Println(fib(30))
    }
}

func fib(n int) int{
    if n < 2 {
        return 1
    }
    return fib(n-1) + fib(n-2)
}
