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
    fib_0, fib_1 := 1, 1
    for ; n >= 1; n-- {
        fib_0, fib_1 = fib_1, fib_0 + fib_1
    }
    return fib_0
}
