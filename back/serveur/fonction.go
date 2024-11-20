package hangmanweb

import (

	hangman "hangmanweb/back/hangman-classique/GAME"
	"strings"

)


var hc = &hangman.HANGMAN{}






func (h *HANGMANWEB) StringTabConv() {
	Data.MotInconnu = strings.Join(hc.MotIconnu, " ")
	
}
