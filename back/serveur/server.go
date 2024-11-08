package hangmanweb

import (
	
	"html/template"
	"net/http"
)



var Data = Donnees{
	Mot: "hc.Mot",
	Find: "C _ I _ _",
	
	

}

var templateHome = "./web/templates/Home.html"
var templatePlay = "./web/templates/Play.html"
var templateWin = "./web/templates/Win.html"
var templateLoose = "./web/templates/Loose.html"

func (h *HANGMANWEB) Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(templateHome))
	tmpl.Execute(w, Data)
	
}

func (h *HANGMANWEB) Play(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(templatePlay))
	tmpl.Execute(w, Data)
}

func (h *HANGMANWEB) Win(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(templateWin))
	tmpl.Execute(w, Data)
}
func (h *HANGMANWEB) Loose(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(templateLoose))
	tmpl.Execute(w, Data)
}