package hangmanweb

import (
	"html/template"
	"net/http"
	"fmt"
)



var Data = Donnees{}

func (h *HANGMANWEB)WebInit() {
	h.StringTabConv()
	
	
	
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
	h.Template(w, r, 0)
	
}

func (h *HANGMANWEB) Play(w http.ResponseWriter, r *http.Request) {
	h.Request(w, r, 1)
	fmt.Println(hc.Mot)
	fmt.Println(hc.Erreur)
	fmt.Println(Data.Usedletter)
}
	

func (h *HANGMANWEB) Win(w http.ResponseWriter, r *http.Request) {
	h.Fin(w, r, 2)
}

func (h *HANGMANWEB) Loose(w http.ResponseWriter, r *http.Request) {
	h.Fin(w, r, 3)
}



func(h *HANGMANWEB) Request(w http.ResponseWriter, r *http.Request, TmplNb int)  {
	r.ParseForm()
	tmpl := template.Must(template.ParseFiles(h.TmplList[TmplNb]))
	lettre := r.Form.Get("lettre")
	h.Lettre = lettre
	h.WebInit()

	if len(h.Lettre) > 0 {
		Data.Usedletter += lettre
		hc.TestLetter(lettre)
		if Data.Essai >= 10 {
			h.Fin(w, r, 3)
		}
	}
	
	tmpl.Execute(w, Data)
	
}
func(h *HANGMANWEB) Fin(w http.ResponseWriter, r *http.Request, TmplNb int)  {
	r.ParseForm()
	tmpl := template.Must(template.ParseFiles(h.TmplList[TmplNb]))
	status := r.Form.Get("restart")
	h.WebInit()

	if status == "restart" {
		tmpl.Execute(w, Data)
	}
	
}
	
	

func(h *HANGMANWEB) Template(w http.ResponseWriter, r *http.Request, TmplNb int)  {
	r.ParseForm()
	tmpl := template.Must(template.ParseFiles(h.TmplList[TmplNb]))
	level := r.Form.Get("level")
	h.Level = level

	if level == "easy" {
		h.WebInit()
		hc.Intiwordlist("/back/hangman-classique/utile/wordlist/words.txt")
		hc.RandomWord()
	} else if level == "medium" {
		h.WebInit()
		hc.Intiwordlist("/back/hangman-classique/utile/wordlist/words2.txt")
		hc.RandomWord()
	
	} else if level == "hard" {
		h.WebInit()
		hc.Intiwordlist("back/hangman-classique/utile/wordlist/words3.txt")
		hc.RandomWord()
	}
	tmpl.Execute(w, Data)
	
	
}
