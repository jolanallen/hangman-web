package hangmanweb

import (
	//"fmt"
	hangman "hangmanweb/back/hangman-classique/GAME"
	"html/template"
	"net/http"
)
var tmpl *template.Template
var hc hangman.HANGMAN

func Home(w http.ResponseWriter,r *http.Request) {
	
}

func Play(w http.ResponseWriter,r *http.Request) {
	//tmpl.Execute(w, "web/templates/Play.html")
}