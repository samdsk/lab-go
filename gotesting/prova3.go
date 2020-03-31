package main

import (
	"encoding/json"
	"esami"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var m esami.Nomi2Voti
	b, _ := ioutil.ReadFile(os.Args[1])
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Fprintf(os.Stderr, "BRUTTO ERRORE: %v\n", err)
		return
	}
	fmt.Println(m)
}
