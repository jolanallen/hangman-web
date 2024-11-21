package hangmanweb

import (
	"bufio"
	"fmt"
	"time"
	"math/rand"
	"os"

)


func (h *HANGMANWEB) WebInit() {
	h.Intiwordlist()
	
	
	h.Erreur = 0
	h.Mot = ""
	h.MotEasy = ""
	h.MotMedium = ""
	h.MotHard = ""
	h.MotIconnu = []string{}
	h.IsRunning = true
	h.YouWin = false
	h.YouLoose =false

	
	h.TemplateList =[]string{"./web/templates/Home.html", "./web/templates/Play.html", "./web/templates/Win.html", "./web/templates/Loose.html"}
	h.LevelPath = []string{"../wordlist/word.txt", "../wordlist/word2.txt", "../wordlist/word3.txt"}
	
	
	
	
}
func (h *HANGMANWEB) Intiwordlist() {
	wordFileEasy, err1 := os.Open(h.LevelPath[0])                   //ouverture des differents fichier texte qui contienent les liste de mots des differents niveau
	wordFileMedium, err2 := os.Open(h.LevelPath[1])  
	wordFileHard, err3 := os.Open(h.LevelPath[2])  

	if err1 != nil || err2 != nil || err3 != nil {                   // si erreur ouverture d'un des fichier alors on les ferme tous et on 
		fmt.Println("Erreur ouverture fichier:", err1, err2, err3)
		wordFileEasy.Close()
		wordFileMedium.Close()
		wordFileHard.Close()
		return
	} else {
		scanner := bufio.NewScanner(wordFileEasy)                   // création de scanners pour scanner chaque fichier texte précédament ouvert 
		scanner2 := bufio.NewScanner(wordFileMedium)
		scanner3 := bufio.NewScanner(wordFileHard)
		for scanner.Scan() { 
			h.WordlsitEasy = append(h.WordlsitEasy, scanner.Text())  
		}
		fmt.Println("WordlistEasy created successfully.")    //message indiquant la création de la liste mot avec succès
		defer wordFileEasy.Close()                          // ferme le fichier texte en dernier grace au defer

		for scanner2.Scan() {
			h.WordlistMedium = append(h.WordlistMedium, scanner2.Text())
		}
		fmt.Println("WordlistMedium created successfully.")
		defer wordFileMedium.Close()
		

		for scanner3.Scan() {
			h.wordlistHard = append(h.wordlistHard, scanner3.Text())
		}
		fmt.Println("WordlistHard created successfully.")
		defer wordFileHard.Close() 
	}

}



  
		


func (h *HANGMANWEB) RandomWord() {
	rand.Seed(time.Now().UnixMilli())  
	
	rdmEasy := rand.Intn(len(h.WordlsitEasy))
	h.MotEasy = h.WordlsitEasy[rdmEasy]

	rdmMedium := rand.Intn(len(h.WordlistMedium))
	h.MotMedium = h.WordlistMedium[rdmMedium]

	rdmHard := rand.Intn(len(h.wordlistHard))
	h.MotHard = h.wordlistHard[rdmHard]	
	fmt.Println(h.MotEasy, h.MotMedium, h.MotHard)
                           
}