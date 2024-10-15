package main

import (
	"fmt"
	//hw "hangmanweb/Hangman-web"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO\n")
}

 func main() {
	http.HandleFunc("/Home", Home)
	http.ListenAndServe(":8080", nil)
 }