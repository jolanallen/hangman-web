package hangmanweb

import (
	
	"encoding/json"
	"fmt"
	hangman "hangmanweb/back/hangman-classique/GAME"
	"io/ioutil"
	"net/http"
	"strings"

)


var hc = &hangman.HANGMAN{}



func (h *HANGMANWEB) WebInit() {
	hc.Erreur = 10
	hc.Intiwordlist()
	hc.RandomWord()
	h.StringTabConv()
	
	
	fmt.Println(hc.Mot)
	fmt.Println(hc.Erreur)
	fmt.Println(Data.Usedletter)
	Data.Mot = hc.Mot
	Data.Essai = hc.Erreur
	Data.Usedletter = "patate"
	
	
	
}

func (h *HANGMANWEB) Testletter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
        // Lire le corps de la requête
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Impossible de lire le corps de la requête", http.StatusBadRequest)
            return
        }

        // Fermer le corps de la requête
        defer r.Body.Close()

        // Décoder les données JSON
        var msg Message
        err = json.Unmarshal(body, &msg)
        if err != nil {
            http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
            return
        }
		fmt.Println(body)
		fmt.Println(msg)

        // Loguer et répondre avec un message de confirmation
        fmt.Printf("Données reçues : %s\n", msg.Message)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status": "Données reçues avec succès"}`))
    } else {
        http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
    }
}






func (h *HANGMANWEB) StringTabConv() {
	Data.Find = strings.Join(hc.MotIconnu, " ")
	Data.Usedletter = strings.Join(hc.UsedLetter, " ")

}
