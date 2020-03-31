package main

import (
  _"fmt"
  "net/http"
  "log"
)

func main() {
  http.HandleFunc("/", page1)

  log.Fatal(http.ListenAndServe(":8080",nil))
}

func page1(w http.ResponseWriter, r *http.Request){
  var text = []byte(`<!DOCTYPE html><head><title>Ciao Mondo</title></head><body><h1>Pagina 1</h1></body>`)

  w.Write(text)
}
