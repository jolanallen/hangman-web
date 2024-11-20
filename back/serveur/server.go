package hangmanweb

import (
	"html/template"
	"net/http"
	"strings"
	
)




var Data = Donnees{}






func (h *HANGMANWEB) Home(w http.ResponseWriter, r *http.Request) {
	h.Request(w, r, 0)
	
}

func (h *HANGMANWEB) Play(w http.ResponseWriter, r *http.Request) {
	h.RequestPlay(w, r, 1)
	
}
	

func (h *HANGMANWEB) Win(w http.ResponseWriter, r *http.Request) {
	h.Request(w, r, 2)
}

func (h *HANGMANWEB) Loose(w http.ResponseWriter, r *http.Request) {
	h.Request(w, r, 3)
}

func(h *HANGMANWEB) Request(w http.ResponseWriter, r *http.Request, nb int)  {
	tmpl := template.Must(template.ParseFiles(h.TemplateList[nb]))
	tmpl.Execute(w, nil)
}

func(h *HANGMANWEB) RequestPlay(w http.ResponseWriter, r *http.Request, Nb int)  {
	
	r.ParseForm()
	tmpl := template.Must(template.ParseFiles(h.TemplateList[Nb]))
	lettre := r.Form.Get("lettre")
	h.Lettre = lettre

	Data := Donnees{
		Mot: h.Mot, 
		Essai: h.Erreur,
		Usedletter: h.Usedletter,
		MotInconnu: strings.Join(h.MotIconnu, " "),
	}
	
	
	tmpl.Execute(w, Data)
	
}


