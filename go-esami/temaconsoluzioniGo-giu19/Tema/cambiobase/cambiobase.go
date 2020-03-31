// per testare da shell 'echo $((11#654))' al reverse, converte in decimale 654 in base 11

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("numero parametri non valido")
		return
	}
	n, err1 := strconv.Atoi(os.Args[1])
	b, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("parametro non valido")
		return
	}

	n_b, err := tenToB(n, b)
	if err == nil {
		fmt.Println(n_b)
	} else {
		fmt.Println(err)
	}
}

func tenToB(n, b int) (base_b string, err error) {
	if b < 2 || b > 16 {
		err = errors.New("base non valida")
		return
	}

	symbols := "0123456789ABCDEF"

	var q, r int
	q, r = n, 0
	for q != 0 {
		r = q % b
		q = q / b
		base_b = string(symbols[r]) + base_b
	}
	return
}
