package hangmanweb

import (
	
	hangman "hangmanweb/back/hangman-classique/GAME"
	"html/template"
	"net/http"
)

var hc *hangman.HANGMAN
var Donnees = Data{
	Essai: 0,
	Mot: "fuck",
	//Essai: hc.Erreur,
	//Mot: hc.Mot,

}

func (h *HANGMANWEB) Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/Home.html"))
	tmpl.Execute(w, Donnees)
	
}

func (h *HANGMANWEB) Play(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/Play.html"))
	tmpl.Execute(w, Donnees)
}

func (h *HANGMANWEB) Win(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/Win.html"))
	tmpl.Execute(w, Donnees)
}
func (h *HANGMANWEB) Loose(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/Loose.html"))
	tmpl.Execute(w, Donnees)
}