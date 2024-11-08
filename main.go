package main

import (
	"fmt"
	hangman "hangmanweb/back/hangman-classique/GAME"
	"hangmanweb/back/serveur"
	"net/http"
)
var hc = &hangman.HANGMAN{}
var web *hangmanweb.HANGMANWEB

func main() {
	web.WebInit()
	
	fmt.Println("Serveur démarré sur http://localhost:8080")
	css := http.FileServer(http.Dir("./web/css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	photo := http.FileServer(http.Dir("./web/utiles"))
	http.Handle("/utiles/", http.StripPrefix("/utiles/", photo))


	// Routes pour les pages de jeu
	http.HandleFunc("/", web.Home)
	http.HandleFunc("/Play", web.Play)
	http.HandleFunc("/Win", web.Win)
	http.HandleFunc("/Loose", web.Loose)
	
	
	// Démarrage du serveur sur le port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur :", err)
	}


}






