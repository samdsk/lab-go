package main

import (
	"encoding/json"
	"esami"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	elenco, err := esami.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERRORE: %v\n", err)
		return
	}

	m := esami.Medie(elenco)
	for nome, media := range m {
		fmt.Printf("%-30s\t%.2f\t%.2f\n", nome, media, esami.Trentesimi2Centodecimi(media))
	}

	b, _ := json.Marshal(elenco)
	ioutil.WriteFile("esami.json", b, 0644)


}
