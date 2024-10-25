package main

import (
	"fmt"
	"hangmanweb/back/serveur" 
	"net/http"
)

func main() {
	fmt.Println("Serveur démarré sur http://localhost:8080")
	web:= &hangmanweb.HANGMANWEB{}
	http.Handle("/utiles/", http.StripPrefix("/utiles/", http.FileServer(http.Dir("/home/jolan/hangman-web/web/utiles"))))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("/home/jolan/hangman-web/web/css"))))



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






