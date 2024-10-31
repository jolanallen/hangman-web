package hangmanweb

import (
	//"fmt"
	"html/template"
	"net/http"
)

var hc *hangman.HANGMAN
var Donnees = Data{
	Essai: 10,
	//Essai: hc.Erreur,
	//Mot: hc.Mot,

}

func (h *HANGMANWEB) Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(templateHome))
	defer
	tmpl.Execute(w, Data)
	
}

func (h *HANGMANWEB) Play(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(templatePlay))
	defer
	tmpl.Execute(w, Data)
}

func (h *HANGMANWEB) Win(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(templateWin))
	defer
	tmpl.Execute(w, Data)
}
func (h *HANGMANWEB) Loose(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(templateLoose))
	defer
	tmpl.Execute(w, Data)
}
func (h *HANGMANWEB) Imput(w http.ResponseWriter, r *http.Request) {
	
}

