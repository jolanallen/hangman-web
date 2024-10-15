package main

import (
	"fmt"
	hangmanweb "hangmanweb/back/serveur"
	"log"
	"net/http"
)

 func main() {
	http.HandleFunc("/Play", hangmanweb.Play)
	http.HandleFunc("/Home", hangmanweb.Home)
	fmt.Println("Serveur démarré sur http://localhost:8080/Home")
	log.Fatal(http.ListenAndServe(":8080", nil))
 }