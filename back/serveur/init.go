package hangmanweb

import (
	"bufio"
	"fmt"
	"time"
	"math/rand"
	"os"

)


func (h *HANGMANWEB) WebInit() {
	rand.Seed(time.Now().UnixNano())
	h.Intiwordlist()
	h.RandomWord()

	h.Usedletter =""
	h.Erreur = 0
	h.MotAdeviner = []string{}
	h.MotIconnu = []string{}

	h.YouWin = false
	h.YouLoose = false


	h.TemplateHome = "./web/templates/Home.html"
	h.TemplatePlay = "./web/templates/Play.html"
	h.TemplateWin = "./web/templates/Win.html"
	h.TemplateLoose = "./web/templates/Loose.html"
	
	
}



func (h *HANGMANWEB) Intiwordlist() {

	wordFileEasy, err1 := os.Open("back/wordlist/words.txt")               //ouverture des differents fichier texte qui contienent les liste de mots des differents niveau

	wordFileMedium, err2 := os.Open("back/wordlist/words2.txt")  
	
	wordFileHard, err3 := os.Open("back/wordlist/words3.txt")  

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


func (h *HANGMANWEB) RandomWord() {            // fonction qui génére un 3 nombres aléatoires qui permettront 	
	rdmEasy := rand.Intn(len(h.WordlsitEasy))
	h.MotEasy = h.WordlsitEasy[rdmEasy]

	rdmMedium := rand.Intn(len(h.WordlistMedium))
	h.MotMedium = h.WordlistMedium[rdmMedium]

	rdmHard := rand.Intn(len(h.wordlistHard))
	h.MotHard = h.wordlistHard[rdmHard]	
	fmt.Println("Mot Aléatoire Easy :",h.MotEasy)
    fmt.Println("Mot Aléatoire medium :", h.MotMedium)    
	fmt.Println("Mot Aléatoire Hard :", h.MotHard)  
	
}




