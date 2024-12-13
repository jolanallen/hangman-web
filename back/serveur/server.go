package hangmanweb

import (
	"html/template"
	"net/http"
	"strings"
)


func (h *HANGMANWEB) Home(w http.ResponseWriter, r *http.Request) {   // fonction pour afficher les differents templates html
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
	niveau := r.Form.Get("level")                                             // gestion des requêtes Get faite par le joueur 
	boutonMenu := r.Form.Get("bouton-menu")
	lettre := r.Form.Get("lettre")

	if boutonMenu == "restart" {
		niveau = h.LevelActual   // permet de garder en mémoire le niveau actuel cette variable n'est pas réintialiser dans la fonction h.webinit()
		h.WebInit()
		
	} else if boutonMenu == "home" {
		h.WebInit()
	}
	
	switch niveau {
	case "easy" :
		h.Mot = h.MotEasy		
		h.InitMotADeviner(h.Mot)       //fonction qui mettra chaque charactére du mot dans un tableau qui sera utiliser pour la comparaisont
		h.LevelActual = "easy"
	case "medium" :
		h.Mot = h.MotMedium
		h.InitMotADeviner(h.Mot)         //fonction qui mettra chaque charactére du mot dans un tableau qui sera utiliser pour la comparaisont
		h.LevelActual = "medium"
	case "hard":
		h.Mot = h.MotHard
		h.InitMotADeviner(h.Mot)         //fonction qui mettra chaque charactére du mot dans un tableau qui sera utiliser pour la comparaisont
		h.LevelActual = "hard"
	}
		h.TestLetter(w, r, lettre)
	
	
	
	Data := Donnees{
		Mot: h.Mot,
		Essai: h.Erreur,
		Usedletter: h.Usedletter,
		MotInconnu: strings.Join(h.MotIconnu, " "),  // ligne un peu particuliere qui permet de convertir un tableau de string en string car html n'affiche pas les tableau 
	}
	
	
	tmpl.Execute(w, Data)                        // eaffichage du templates avec les données actualiser dans la structure data
	

}





