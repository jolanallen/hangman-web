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

	if boutonMenu == "restart" {
		niveau = h.LevelActual
		h.WebInit()
		
	}
		
	if niveau == "easy" {
		h.Mot = h.MotEasy		
		h.InitMotaDeviner(h.Mot)       //fonction qui mettra chaque charactére du mot dans un tableau qui sera utiliser pour la comparaisont
		h.LevelActual = "easy"

	} else if niveau == "medium" {
		h.Mot = h.MotMedium
		h.InitMotaDeviner(h.Mot)         //fonction qui mettra chaque charactére du mot dans un tableau qui sera utiliser pour la comparaisont
		h.LevelActual = "medium"

	} else if niveau == "hard" {
		h.Mot = h.MotHard
		h.InitMotaDeviner(h.Mot)         //fonction qui mettra chaque charactére du mot dans un tableau qui sera utiliser pour la comparaisont
		h.LevelActual = "hard"
	}

	if !(strings.Contains(h.Usedletter, lettre)) {  // vérifie si la lettre a déja été utiliser 
		h.Usedletter += lettre                             // si c'est pas le cas elle ajoute la lettre au letre utilisé 
		if strings.Contains(h.Mot,  lettre) {
			h.TestLetter(lettre)
		
		} else if h.Erreur <= 9 {
			h.Erreur += 1
		
		}  else {
			http.Redirect(w, r, "/Loose", http.StatusSeeOther)
		}
		
	}

	// for i := 0; i <= len(h.MotAdeviner); i++ {
	// 	if h.MotAdeviner[i] == h.MotIconnu[i] {
	// 		http.Redirect(w, r, "/Win", http.StatusSeeOther)
	// 	}
	// }

	
	
	Data := Donnees{
		Mot: h.Mot, 
		Essai: h.Erreur,
		Usedletter: h.Usedletter,
		MotInconnu: strings.Join(h.MotIconnu, ""),
	}
	
	
	tmpl.Execute(w, Data)
	

}





