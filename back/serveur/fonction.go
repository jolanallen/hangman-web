package hangmanweb

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	
)



func (h *HANGMANWEB) InitMotADeviner(Mot string) {
	for _, char := range h.Mot {
    h.MotAdeviner = append(h.MotAdeviner, string(char))
  }
    
  for i :=  0; i < len(h.Mot); i++ {                /// crée un mot inconnu qui contient uniquement des "_" qui corresponde a chaque caractéres 
		h.MotIconnu = append(h.MotIconnu, "_")  
	}
  
	rdmNB := rand.Intn(len(h.Mot))                 // permet de révéler aléatoirement une lettre dans le mot 
	h.MotIconnu[rdmNB] = h.MotAdeviner[rdmNB]
  fmt.Println(h.MotIconnu)
}   




func (h *HANGMANWEB) TestLetter(w http.ResponseWriter, r *http.Request, lettre string) {

  if !(strings.Contains(h.Usedletter, lettre)) {  // vérifie si la lettre a déja été utiliser 
		h.Usedletter += lettre                             // si c'est pas le cas elle ajoute la lettre au letre utilisé 
		if strings.Contains(h.Mot,  lettre) {
      for i := 0; i < len(h.MotAdeviner); i++ {        // boucle qui vérifie si la lettre entrée par le joueur appartient au mot a devienr si oui il l'ajout au mot inconnu
        if h.MotAdeviner[i] == lettre {
            h.MotIconnu[i] =  lettre
            fmt.Println(h.MotIconnu)
        }
      }
			h.TestWord(w, r)                              //fonction qui test si le mot à était deviner 
		} else if h.Erreur < 10 {
			h.Erreur += 1
      if h.Erreur == 10 {
        http.Redirect(w, r, "/Loose", http.StatusSeeOther)            // crée une redirection vers la pages Loose 
      }
		} 
	}
  
}
func (h *HANGMANWEB) TestWord(w http.ResponseWriter, r *http.Request) {
  TestMot := strings.Join(h.MotIconnu, "")
  if !(strings.Contains(TestMot, "_")) {                              /// verifie si le mot contient des "_" si ce n'est pas le cas la fonction s'arrete car cela sigifie que le mot est incomplet
    if TestMot == h.Mot {
      http.Redirect(w, r, "/Win", http.StatusSeeOther)                // créée une redirection ver la pages win 
    }
  } 
}




