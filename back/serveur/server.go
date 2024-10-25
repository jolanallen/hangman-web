package hangmanweb

import (
	"fmt"
	"html/template"
	"net/http"
)



// Méthode pour afficher une page HTML en chargeant un template spécifique
func (h *HANGMANWEB) AfficherPageHTML(rw http.ResponseWriter, cheminTemplate string, donnees interface{}) {
	// Charger le fichier template HTML
	pageHTML, err := template.ParseFiles(cheminTemplate)
	if err != nil {
		fmt.Println("Erreur chargement template :", err)
		http.Error(rw, "Erreur lors du chargement de la page", http.StatusInternalServerError)
		return
	}

	// Exécuter le template avec les données fournies
	err = pageHTML.Execute(rw, donnees)
	if err != nil {
		fmt.Println("Erreur execution template :", err)
		http.Error(rw, "Erreur lors de l'affichage de la page", http.StatusInternalServerError)
	}
}

// Gestionnaire pour la page d'accueil
func (h *HANGMANWEB) Home(rw http.ResponseWriter, r *http.Request) {
	h.AfficherPageHTML(rw, "./web/templates/Home.html", h.Donnees)
}

// Gestionnaire pour la page de jeu
func (h *HANGMANWEB) Play(rw http.ResponseWriter, r *http.Request) {
	h.AfficherPageHTML(rw, "./web/templates/Play.html", h.Donnees)
}

// Gestionnaire pour la page de défaite
func (h *HANGMANWEB) Loose(rw http.ResponseWriter, r *http.Request) {
	h.AfficherPageHTML(rw, "./web/templates/Loose.html", h.Donnees)
}

// Gestionnaire pour la page de victoire
func (h *HANGMANWEB) Win(rw http.ResponseWriter, r *http.Request) {
	h.AfficherPageHTML(rw, "./web/templates/Win.html", h.Donnees)
}
