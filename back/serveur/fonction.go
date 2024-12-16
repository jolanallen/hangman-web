package hangmanweb

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	
)



func (h *HANGMANWEB) InitSecretWord(Word string) {
	for _, char := range h.Word {
    h.SecretWord = append(h.SecretWord, string(char))
  }
    
  for i :=  0; i < len(h.Word); i++ {                /// crée un Word inconnu qui contient uniquement des "_" qui corresponde a chaque caractéres 
		h.HiddenWord = append(h.HiddenWord, "_")  
	}
  
	rdmNB := rand.Intn(len(h.Word))                 // permet de révéler aléatoirement une letter dans le Word 
	h.HiddenWord[rdmNB] = h.SecretWord[rdmNB]
  fmt.Println(h.HiddenWord)
}   




func (h *HANGMANWEB) TestLetter(w http.ResponseWriter, r *http.Request, letter string) {

  if !(strings.Contains(h.Usedletter, letter)) {  // vérifie si la letter a déja été utiliser 
		h.Usedletter += letter
    h.Usedletter += " "                             // si c'est pas le cas elle ajoute la letter au letre utilisé 
		if strings.Contains(h.Word,  letter) {
      for i := 0; i < len(h.SecretWord); i++ {        // boucle qui vérifie si la letter entrée par le joueur appartient au Word a devienr si oui il l'ajout au Word inconnu
        if h.SecretWord[i] == letter {
            h.HiddenWord[i] =  letter
            fmt.Println(h.HiddenWord)
        }
      }
			h.TestWord(w, r)                              //fonction qui test si le Word à était deviner 
		} else if h.Error < 10 {
			h.Error += 1
      if h.Error == 10 {
        http.Redirect(w, r, "/Loose", http.StatusSeeOther)            // crée une redirection vers la pages Loose 
      }
		} 
	}
  
}
func (h *HANGMANWEB) TestWord(w http.ResponseWriter, r *http.Request) {
  TestWord := strings.Join(h.HiddenWord, "")
  if !(strings.Contains(TestWord, "_")) {                              /// verifie si le Word contient des "_" si ce n'est pas le cas la fonction s'arrete car cela sigifie que le Word est incomplet
    if TestWord == h.Word {
      http.Redirect(w, r, "/Win", http.StatusSeeOther)                // créée une redirection ver la pages win 
    }
  } 
}




