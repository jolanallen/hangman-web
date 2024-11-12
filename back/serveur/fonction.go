package hangmanweb

import (
	"fmt"
	hangman "hangmanweb/back/hangman-classique/GAME"
	"strings"
)


var hc = &hangman.HANGMAN{}



func (h *HANGMANWEB) WebInit() {
	hc.Intiwordlist()
	hc.RandomWord()
	h.StringTabConv()
	
	
	fmt.Println(hc.Mot)
	fmt.Println(hc.Erreur)
	fmt.Println(Data.Usedletter)
	Data.Mot = hc.Mot
	Data.Essai = hc.Erreur
	
	
	
}

func (h* HANGMANWEB) Testletter() {

}




func (h *HANGMANWEB) StringTabConv() {
	Data.Find = strings.Join(hc.MotIconnu, " ")
	Data.Usedletter = strings.Join(hc.UsedLetter, " ")

}
