package hangmanweb

import (
	"html/template"
	"net/http"
	"fmt"
)



var Data = Donnees{}

func (h *HANGMANWEB)WebInit() {
	hc.Intiwordlist()
	hc.RandomWord()
	h.StringTabConv()
	
	fmt.Println(hc.Mot)
	fmt.Println(hc.Erreur)
	fmt.Println(Data.Usedletter)
	
	Data.Mot = hc.Mot
	Data.Essai = hc.Erreur
	Data.Usedletter = ""

	h.TemplateLoose ="./web/templates/Loose.html"
	h.TemplateHome = "./web/templates/Home.html"
	h.TemplatePlay = "./web/templates/Play.html"
	h.TemplateWin = "./web/templates/Win.html"

	h.TmplList = []string{h.TemplateHome, h.TemplatePlay, h.TemplateWin, h.TemplateLoose}
	
	
	
}




func (h *HANGMANWEB) Home(w http.ResponseWriter, r *http.Request) {
	h.RequestAndTemplate(w, r, 0)
	
}

func (h *HANGMANWEB) Play(w http.ResponseWriter, r *http.Request) {
	h.RequestAndTemplate(w, r, 1)
}
	

func (h *HANGMANWEB) Win(w http.ResponseWriter, r *http.Request) {
	h.RequestAndTemplate(w, r, 2)
}

func (h *HANGMANWEB) Loose(w http.ResponseWriter, r *http.Request) {
	h.RequestAndTemplate(w, r, 3)
}



func(h *HANGMANWEB) RequestAndTemplate(w http.ResponseWriter, r *http.Request, TmplNb int)  {
	r.ParseForm()
	tmpl := template.Must(template.ParseFiles(h.TmplList[TmplNb]))
	lettre := r.Form.Get("lettre")
	level := r.Form.Get("level")
	h.Lettre = lettre
	h.Level = level
	
	
}
