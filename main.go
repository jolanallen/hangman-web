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
	fmt.Println("Serveur démarré sur http://localhost:8082")
	css := http.FileServer(http.Dir("./web/css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	photo := http.FileServer(http.Dir("./web/utiles"))
	http.Handle("/utiles/", http.StripPrefix("/utiles/", photo))
	js := http.FileServer(http.Dir("./web/js"))
	http.Handle("/js/", http.StripPrefix("/js/", js))


	// Routes pour les pages de jeu
	http.HandleFunc("/", web.Home)
	http.HandleFunc("/hangman", web.Play)
	http.HandleFunc("/Win", web.Win)
	http.HandleFunc("/Loose", web.Loose)
	

	// Démarrage du serveur sur le port 8080
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur :", err)
	}


}






