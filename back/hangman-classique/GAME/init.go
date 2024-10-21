package hangman

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"math/rand"
	"time"
)



func (hangman *HANGMAN) Init() {
		hangman.drawDisplay()
		hangman.intiwordlist()
		hangman.initHangman()
		hangman.randomWord()
}
func (hangman *HANGMAN) Flag() {
	help := flag.Bool("h", false, "Afficher de l'aide")  /// flag pour afficher l'aide
	hard := flag.Bool("hard", false, "mode de jeu hard mots difficiles")  /// flag pour choisir le mode de difficulté difficile
	medium := flag.Bool("medium", false, "mode de jeu medium mots moyennement difficiles") /// flag pour choisir le mode de difficulté intermédiaire

	flag.Parse() /// Analyse les flags passés en ligne de commande.

	if *help {
		flag.Usage()  /// Affiche l'aide si l'utilisateur a choisi l'option `-h`.
	}
	if *hard {
		hangman.hard = true      /// Active le mode difficile si `-hard` est passé.
		fmt.Println("mode de jeu hard")
	} else if *medium {
		hangman.medium = true    /// Active le mode intermédiaire si `-medium` est passé.
		fmt.Println("mode de jeu medium")
	}
}

func (hangman *HANGMAN) drawDisplay() {
	fmt.Println("Bienvenue dans le HANGMAN Ynov 2024 B1 Info Montpellier )")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println((`																													 
					 __ __   ____  ____    ____  ___ ___   ____  ____  				\ / |\ | / \ \  /
   					|  |  | /    ||    \  /    ||   |   | /    ||    \ 				 |  | \| \_/  \/
   					|  |  ||  o  ||  _  ||   __|| _   _ ||  o  ||  _  |
   					|  _  ||     ||  |  ||  |  ||  \_/  ||     ||  |  |
   					|  |  ||  _  ||  |  ||  |_ ||   |   ||  _  ||  |  |
   					|  |  ||  |  ||  |  ||     ||   |   ||  |  ||  |  |
   					|__|__||__|__||__|__||___,_||___|___||__|__||__|__|
	   `))
	fmt.Println("                                                   	  ")
	fmt.Println("					                     Bienvenue !						  ")
	fmt.Println("                                                   	  ")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------------")
}

func (hangman *HANGMAN) intiwordlist() {
	if hangman.medium {
		hangman.wordFile = "utile/wordlist/words2.txt"  /// Utilise une liste de mots intermédiaires pour le mode médium.
		fmt.Println("mode de jeu : intermédiaire")
	} else if hangman.hard {
		hangman.wordFile = "utile/wordlist/words3.txt"  /// Utilise une liste de mots difficiles pour le mode hard.
		fmt.Println("mode de jeu : hard")
	} else {
		hangman.wordFile = "utile/wordlist/words.txt"   /// Utilise la liste de mots par défaut (facile).
		fmt.Println("mode de jeu : soft (défaut)")
		fmt.Println("tapez -medium pour un mode de jeu intermédiaire")
		fmt.Println("tapez -hard pour un mode de jeu difficile")
	}
		
	wordFile, err := os.Open(hangman.wordFile)  /// Ouvre le fichier de la liste de mots.
	if err != nil {
		fmt.Println("Erreur ouverture fichier:", err)  /// Affiche une erreur si le fichier ne peut être ouvert.
		defer wordFile.Close()
	}
	if err == nil {
		scanner := bufio.NewScanner(wordFile)
		for scanner.Scan() { 
			hangman.TabMots = append(hangman.TabMots, scanner.Text())  /// Ajoute chaque mot du fichier à la liste `TabMots`.
		}
		defer wordFile.Close()  /// Ferme le fichier après lecture.
	}
}

func (hangman *HANGMAN) randomWord() {
	rand.Seed(time.Now().UnixMilli())               /// Utilise la date du jour  pour générer un nombre aléatoire unixmili correspond a la date de reférence de Unix 01/01/1970
	hangman.randomNb = rand.Intn(len(hangman.TabMots))                
	hangman.Mot = hangman.TabMots[hangman.randomNb]             /////// Sélectionne le mot à deviner à partir du nombre aléatoire dans le tableau de mots 
	
	for _, char := range hangman.Mot {
		hangman.MotAdeviner = append(hangman.MotAdeviner, string(char))  /// Ajoute chaque lettre du mot à deviner dans `MotAdeviner`.
	}
	
	for i :=  0; i < len(hangman.Mot); i++ {
		hangman.motIconnu = append(hangman.motIconnu, "_")  ///////::// Remplace chaque lettre par un underscore pour masquer le mot
	}
	
	hangman.randomNb = rand.Intn(len(hangman.Mot))                    ///://////:: Sélectionne une lettre aléatoire à dévoiler
	hangman.motIconnu[hangman.randomNb] = hangman.MotAdeviner[hangman.randomNb]              /// Révèle une lettre du mot à deviner
	hangman.IsRunning = true                     /// Indique que le jeu est en cours.
}

func (hangman *HANGMAN) initHangman() {
	hangman.erreur = 0                                         /// Initialise le compteur d'erreurs à 0.

	hangFile, err := os.Open("utile/hangman/hangman.txt")  /// Ouvre le fichier contenant les états du pendu.
	if err != nil {
		fmt.Println(err)
		defer hangFile.Close()
	}
  	if err == nil {
		scanner := bufio.NewScanner(hangFile)
    	
		for scanner.Scan() { 
     		if scanner.Text() == "" {
       			hangman.TabHang = append(hangman.TabHang, hangman.etat.String() )  /// Ajoute l'état graphique du pendu lorsqu'une ligne vide est rencontrée.
       			hangman.etat.Reset()
     		} else {
        		hangman.etat.WriteString(scanner.Text() + "\n")  /// Continue à lire les lignes pour ajouter l'état du pendu.
     		}
		}
		defer hangFile.Close()  /// Ferme le fichier après lecture.
	}
}