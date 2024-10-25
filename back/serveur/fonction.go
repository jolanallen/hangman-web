package hangmanweb
import (
	"fmt"
	"html/template"
	"net/http"
)
func (hangmanweb *HANGMANWEB) AfficherPageHTML(rw http.ResponseWriter, cheminTemplate string, Donnees interface{} ) {
	// Charger le fichier template HTML
	
	pageHTML, err := template.ParseFiles(cheminTemplate)
	if err == nil {
		err = pageHTML.Execute(rw, Donnees)
		if err != nil {
			fmt.Println("erreur execution template")

		}
		
	} else {
		fmt.Println("erreur chargement template")
	}
	
}