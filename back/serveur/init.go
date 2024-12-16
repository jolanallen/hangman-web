package hangmanweb

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

)


func (h *HANGMANWEB) WebInit() {
	h.RandomWord()

	h.Usedletter =""
	h.Error = 0
	h.SecretWord = []string{}
	h.HiddenWord = []string{}

	h.TemplateHome = "./web/templates/Home.html"
	h.TemplatePlay = "./web/templates/Play.html"
	h.TemplateWin = "./web/templates/Win.html"
	h.TemplateLoose = "./web/templates/Loose.html"
	
	
}



func (h *HANGMANWEB) Intiwordlist() {

	wordFileEasy, err1 := os.Open("back/wordlist/words.txt")               //ouverture des differents fichier texte qui contienent les liste de Words des differents niveau

	wordFileMedium, err2 := os.Open("back/wordlist/words2.txt")  
	
	wordFileHard, err3 := os.Open("back/wordlist/words3.txt")  

	if err1 != nil || err2 != nil || err3 != nil {                   // si Error ouverture d'un des fichier alors on les ferme tous et on 
		fmt.Println("Error ouverture fichier:", err1, err2, err3)
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
		fmt.Println("WordlistEasy create successfully.")    //message indiquant la création de la liste Word avec succès
		defer wordFileEasy.Close()                          // ferme le fichier texte en dernier grace au defer

		for scanner2.Scan() {
			h.WordlistMedium = append(h.WordlistMedium, scanner2.Text())
		}
		fmt.Println("WordlistMedium create successfully.")
		defer wordFileMedium.Close()
		

		for scanner3.Scan() {
			h.wordlistHard = append(h.wordlistHard, scanner3.Text())
		}
		fmt.Println("WordlistHard create successfully.")
		defer wordFileHard.Close() 
	}

}


func (h *HANGMANWEB) RandomWord() {            		// fonction qui génére un 3 nombres aléatoires qui permettront 	
	rdmEasy := rand.Intn(len(h.WordlsitEasy))	// Générer un nombre aléatoire entre 0 et la longeur de worlist easy 
	h.WordEasy = h.WordlsitEasy[rdmEasy] 
	fmt.Println("easy word was created succesfully")

	rdmMedium := rand.Intn(len(h.WordlistMedium))
	h.WordMedium = h.WordlistMedium[rdmMedium]
	fmt.Println("medium word was created succesfully")

	rdmHard := rand.Intn(len(h.wordlistHard))
	h.WordHard = h.wordlistHard[rdmHard]	
	fmt.Println("hard word was created succesfully")
	
}




