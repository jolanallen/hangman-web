package main

import (
	"fmt"
	"hangmanweb/back/serveur"
	"net/http"
)

var web = &hangmanweb.HANGMANWEB{}
//var web *hangmanweb.HANGMANWEB


func main() {
	web.WebInit()
	fmt.Println("Serveur démarré sur http://localhost:3030")
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
	
	
	
	// Démarrage du serveur sur le port 3030
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur :", err)
	}


}






