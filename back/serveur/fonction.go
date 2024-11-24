package hangmanweb

import "fmt"


func (h *HANGMANWEB) Run() {
	h.WebInit()
	
}



func (h *HANGMANWEB) InitMotaDeviner(Mot string) {
	for _, char := range h.Mot {
        h.MotAdeviner = append(h.MotAdeviner, string(char))
    }
    fmt.Println(h.MotAdeviner)
    // for i := 0; i <= len(h.MotAdeviner); i++ {
    //     h.MotIconnu[i] = h.MotAdeviner[i]
        
    // }
    // fmt.Println(h.MotIconnu) 
}   




func (h *HANGMANWEB) TestLetter(lettre string) {


}




