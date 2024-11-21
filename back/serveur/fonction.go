package hangmanweb

import (
	
	
)


func (h *HANGMANWEB) Run() {
	h.WebInit()
	for h.IsRunning {
        h.YouWin = true
    }
}








func (h *HANGMANWEB) TestLetter(lettre string) {
    h.lettreIsGood = false
    
    for l := range h.MotAdeviner {
        if lettre == string(h.MotAdeviner[l]) {
            h.lettreIsGood = true
            h.MotIconnu[l] = h.MotAdeviner[l] // Utilise la lettre originale pour l'affichage
        }
    }
    if !h.lettreIsGood {
        h.Erreur++
    }
}




