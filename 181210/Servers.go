package main

import(
	"fmt"
	"net/http"
	"html"
	"log"
	)


func fooHandler(w http.ResponseWriter, r *http.Request){}

func main(){
 http.HandleFunc("/foo", fooHandler)

 http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
 })

 log.Fatal(http.ListenAndServe(":8080", nil))
}
