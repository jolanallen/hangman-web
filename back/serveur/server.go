package hangmanweb

import (
	"fmt"
	"net/http"
)


func (hangmanweb *HANGMANWEB) Home(rw http.ResponseWriter, r *http.Request) {
	hangmanweb.AfficherPageHTML(rw, "./web/templates/Home.html", hangmanweb.Donnees)
}

func (hangmanweb *HANGMANWEB) Play(rw http.ResponseWriter, r *http.Request) {
	hangmanweb.AfficherPageHTML(rw, "./web/templates/play.html", hangmanweb.Donnees)
}

func (hangmanweb *HANGMANWEB) Loose(rw http.ResponseWriter, r *http.Request) {
	hangmanweb.AfficherPageHTML(rw, "./web/templates/Loose.html", hangmanweb.Donnees)
}

func (hangmanweb *HANGMANWEB) Win(rw http.ResponseWriter, r *http.Request) {
	hangmanweb.AfficherPageHTML(rw, "./web/templates/Win.html", hangmanweb.Donnees)
}

func (hangmanweb *HANGMANWEB) Serveur() {
	fmt.Println("(http://localhost:8080)")
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("css"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("utiles"))))



	http.HandleFunc("/", hangmanweb.Home)
	http.HandleFunc("/Play", hangmanweb.Play)
	http.HandleFunc("/Win", hangmanweb.Win)
	http.HandleFunc("/Loose", hangmanweb.Loose)

	
	http.ListenAndServe(":8080", nil)
}