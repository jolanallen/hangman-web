package hangmanweb

import (
	"fmt"
	hangman "hangmanweb/back/hangman-classique/GAME"
	"strings"

)


var hc = &hangman.HANGMAN{}
var Data = Donnees{}

func (h *HANGMANWEB)WebInit() {
	hc.Intiwordlist()
	hc.RandomWord()
	h.StringTabConv()
	
	fmt.Println(hc.Mot)
	fmt.Println(hc.Erreur)
	fmt.Println(Data.Usedletter)
	
	Data.Mot = hc.Mot
	Data.Essai = hc.Erreur
	Data.Usedletter = ""

	h.TemplateLoose ="./web/templates/Loose.html"
	h.TemplateHome = "./web/templates/Home.html"
	h.TemplatePlay = "./web/templates/Play.html"
	h.TemplateWin = "./web/templates/Win.html"

	h.TmplList = []string{h.TemplateHome, h.TemplatePlay, h.TemplateWin, h.TemplateLoose}
	
	
	
}






func (h *HANGMANWEB) StringTabConv() {
	Data.MotInconnu = strings.Join(hc.MotIconnu, " ")
	
}
