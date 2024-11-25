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
  fmt.Println(h.MotAdeviner)
    
   
  for i :=  0; i < len(h.Mot); i++ {
		h.MotIconnu = append(h.MotIconnu, "_")  
	}
  
	rdmNB := rand.Intn(len(h.Mot))                 
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
func (h *HANGMANWEB) TestWord(w http.ResponseWriter, r *http.Request) {
  TestMot := strings.Join(h.MotIconnu, "")
  if !(strings.Contains(TestMot, "_")) {
    if TestMot == h.Mot {
      http.Redirect(w, r, "/Win", http.StatusSeeOther)
    }
  } 
}




