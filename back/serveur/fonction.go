package hangmanweb

import (
	"fmt"
	"math/rand"
	"time"
)


func (h *HANGMANWEB) Run() {
	h.WebInit()
	
}



func (h *HANGMANWEB) InitMotaDeviner(Mot string) {
	for _, char := range h.Mot {
        h.MotAdeviner = append(h.MotAdeviner, string(char))
    }
    fmt.Println(h.MotAdeviner)
    
   
    for i :=  0; i < len(h.Mot); i++ {
		h.MotIconnu = append(h.MotIconnu, "_")  
	}
    rand.Seed(time.Now().UnixMilli()) 
	rdmNB:= rand.Intn(len(h.MotAdeviner))                  
	h.MotIconnu[rdmNB] = h.MotAdeviner[rdmNB]
    fmt.Println(h.MotIconnu)
}   




func (h *HANGMANWEB) TestLetter(lettre string) {
  for i := 0; i < len(h.MotAdeviner); i++ {
    if h.MotAdeviner[i] == lettre {
        h.MotIconnu[i] =  lettre
        fmt.Println(h.MotIconnu)
    }
  }
}




