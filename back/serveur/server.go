package hangmanweb

import (
	"html/template"
	"net/http"
	"strings"
	
)











func (h *HANGMANWEB) Home(w http.ResponseWriter, r *http.Request) {
	h.Request(w, r, h.TemplateHome)
	
}

func (h *HANGMANWEB) Play(w http.ResponseWriter, r *http.Request) {
	h.Request(w, r, h.TemplatePlay)
	
}

func (h *HANGMANWEB) Win(w http.ResponseWriter, r *http.Request) {
	h.Request(w, r, h.TemplateWin)
}

func (h *HANGMANWEB) Loose(w http.ResponseWriter, r *http.Request) {
	h.Request(w, r, h.TemplateLoose)
}

func(h *HANGMANWEB) Request(w http.ResponseWriter, r *http.Request, html string)  {
	r.ParseForm()
	tmpl := template.Must(template.ParseFiles(html))
	niveau := r.Form.Get("level")
	h.Level = niveau
	boutonMenu := r.Form.Get("bouton-menu")
	h.Commande = boutonMenu
	lettre := r.Form.Get("lettre")
	h.Lettre = lettre

	if lettre <= "z" && lettre >= "a" {
		if strings.Contains(h.Usedletter, lettre) == false {
			h.Usedletter += lettre
		}

		
		
	}

	if niveau == "easy" {
		h.Mot = h.MotEasy
		h.Level = "easy"
		h.InitMotInconnu(h.Mot)
	} else if niveau == "medium" {
		h.Mot = h.MotMedium
		h.Level = "medium"
		h.InitMotInconnu(h.Mot)
	} else if niveau == "hard" {
		h.Mot = h.MotHard
		h.Level = "hard"
		h.InitMotInconnu(h.Mot)
	}


	
	Data := Donnees{
		Mot: h.Mot, 
		Essai: h.Erreur,
		Usedletter: h.Usedletter,
		MotInconnu: strings.Join(h.MotIconnu, " "),
		Level: h.Level,
	}
	
	
	tmpl.Execute(w, Data)
	

	
}





