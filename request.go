package hangmanweb

import (
	hangman "hangmanweb/hangman-classique/GAME"
	"net/http"
	"html/template"
)
var templates *template.Template
var hc hangman.HANGMAN

func Home(w http.ResponseWriter,r *http.Request) {
	templates.ExecuteTemplate(w, "Home.html", nil)
}
func PLAY(w http.ResponseWriter,r *http.Request) {
	templates.ExecuteTemplate(w, "Play.html", nil)
}