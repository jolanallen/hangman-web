package hangmanweb

import (
	
	hangman "hangmanweb/back/hangman-classique/GAME"
	"html/template"
	"net/http"
)

var hc *hangman.HANGMAN
var Donnees = Data{
	//Essai: hc.Erreur,
	Mot: "fuck",

}

func (h *HANGMANWEB) Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/Home.html"))
	tmpl.Execute(w, Donnees)
	
}

func (h *HANGMANWEB) Play(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/Play.html"))
	tmpl.Execute(w, Donnees)
}

func (h *HANGMANWEB) Win(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/Win.html"))
	tmpl.Execute(w, Donnees)
}
func (h *HANGMANWEB) Loose(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/Loose.html"))
	tmpl.Execute(w, Donnees)
}
// // Méthode pour afficher une page HTML en chargeant un template spécifique
// func (h *HANGMANWEB) AfficherPageHTML(rw http.ResponseWriter, cheminTemplate string, donnees interface{}) {
	
// 	// Charger le fichier template HTML
// 	pageHTML, err := template.ParseFiles(cheminTemplate)
// 	if err != nil {
// 		fmt.Println("Erreur chargement template :", err)
// 		http.Error(rw, "Erreur lors du chargement de la page", http.StatusInternalServerError)
// 		return
// 	}
// 	// Exécuter le template avec les données fournies
// 	err = pageHTML.Execute(rw, Donnees)
// 	if err != nil {
// 		fmt.Println("Erreur execution template :", err)
// 		http.Error(rw, "Erreur lors de l'affichage de la page", http.StatusInternalServerError)
// 	}
	
// }

// Gestionnaire pour la page d'accueil
// func (h *HANGMANWEB) Home(rw http.ResponseWriter, r *http.Request) {
// 	h.AfficherPageHTML(rw, "./web/templates/Home.html", Donnees)
// }

// // Gestionnaire pour la page de jeu
// func (h *HANGMANWEB) Play(rw http.ResponseWriter, r *http.Request) {
// 	h.AfficherPageHTML(rw, "./web/templates/Play.html", Donnees)
	
// }
// // Gestionnaire pour la page de défaite
// func (h *HANGMANWEB) Loose(rw http.ResponseWriter, r *http.Request) {
// 	h.AfficherPageHTML(rw, "./web/templates/Loose.html", h.Donnees)
// }

// // Gestionnaire pour la page de victoire
// func (h *HANGMANWEB) Win(rw http.ResponseWriter, r *http.Request) {
// 	h.AfficherPageHTML(rw, "./web/templates/Win.html", h.Donnees)
// }
