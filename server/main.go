package main

import (
	"fmt"
	//hw "hangmanweb/Hangman-web"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HOME\n")
}
func Jeu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "JEU")
}


 func main() {
	http.HandleFunc("/Home", Home)
	http.HandleFunc("/Jeu", Jeu)
	http.ListenAndServe(":8080", nil)
 }