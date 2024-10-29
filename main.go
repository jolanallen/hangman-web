package main

import (
	"fmt"
	hangman "hangmanweb/back/hangman-classique/GAME"
	"hangmanweb/back/serveur"
	"net/http"
)
var hc *hangman.HANGMAN
var web *hangmanweb.HANGMANWEB

func main() {
	web.Mot()
	fmt.Println("Serveur démarré sur http://localhost:8080")
	
	http.Handle("/utiles/", http.StripPrefix("/utiles/", http.FileServer(http.Dir("/web/utiles"))))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("/web/css"))))

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






