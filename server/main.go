package main

import (
	"fmt"
	"hangmanweb"
	"log"
	"net/http"
	//"html/template"
	
)

 func main() {
	http.HandleFunc("/Play", hangmanweb.PLAY)
	http.HandleFunc("/Home", hangmanweb.Home)
	fmt.Println("Serveur démarré sur http://localhost:8080/Home")
	log.Fatal(http.ListenAndServe(":8080", nil))
 }