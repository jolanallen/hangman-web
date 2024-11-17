package hangman

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "unicode"
)


// normalize convertit une chaîne en enlevant les accents et en la mettant en minuscules
func normalize(s string) string {
    t := ""
    for _, c := range s {
        switch c {
        case 'à', 'â', 'ä':
            t += "a"
        case 'é', 'è', 'ê', 'ë':
            t += "e"
        case 'î', 'ï':
            t += "i"
        case 'ô', 'ö':
            t += "o"
        case 'ù', 'û', 'ü':
            t += "u"
        case 'ç':
            t += "c"
        default:
            t += string(unicode.ToLower(c))
        }
    }
    return t
}

func (hangman *HANGMAN) Run() {
    hangman.wordIsGood = false
    for hangman.IsRunning {
        hangman.Hangman()
        hangman.readletter()
        hangman.testword()
    }
}

func (hangman *HANGMAN) readletter() {
    fmt.Println(hangman.MotIconnu)
    var Reader = bufio.NewReader(os.Stdin)
    String, _ := Reader.ReadString('\n')
    String = strings.TrimSpace(String)
    String = normalize(String) // Normalise l'entrée utilisateur (enlève les accents)

    if String >= "a" && String <= "z" {
        if len(String) > 1 {
            if String == normalize(strings.Join(hangman.MotAdeviner, "")) { // Compare avec le mot normalisé
               // hangman.win()
            } else {
                hangman.Erreur += 2
            }
        } else if len(String) == 1 {
            hangman.lettre = String
            hangman.UsedLetter = append(hangman.UsedLetter, hangman.lettre)
            hangman.testLetter()
        }
    } else {
        fmt.Println("lettre incorrecte !!")
    }
}

func (hangman *HANGMAN) testLetter() {
    hangman.lettreIsGood = false
    normalizedMotAdeviner := normalize(strings.Join(hangman.MotAdeviner, ""))
    for l := range normalizedMotAdeviner {
        if hangman.lettre == string(normalizedMotAdeviner[l]) {
            hangman.lettreIsGood = true
            hangman.MotIconnu[l] = hangman.MotAdeviner[l] // Utilise la lettre originale pour l'affichage
        }
    }
    if !hangman.lettreIsGood {
        hangman.Erreur++
    }
    if hangman.Erreur >= 9 {
        //hangman.gameOver()
    }
}

func (hangman *HANGMAN) testword() {
    hangman.wordIsGood = true
    for l := range hangman.MotIconnu {
        if hangman.MotIconnu[l] == "_" {
            hangman.wordIsGood = false
            break
        }
    }
    // if hangman.wordIsGood {
    //     //hangman.win()
    // }
    // if hangman.Erreur >= 10 {
    //    // hangman.gameOver()
    // }
}


// func (hangman *HANGMAN) win() {
// 		fmt.Println(`
//  		__     __          __          ___       _ 
//  		\ \   / /          \ \        / (_)     | |
// 	 	 \ \_/ /__  _   _   \ \  /\  / / _ _ __ | |
//  	 	  \   / _ \| | | |   \ \/  \/ / | | '_ \| |
//   		   | | (_) | |_| |    \  /\  /  | | | | |_|
//   	 	   |_|\___/ \__,_|     \/  \/   |_|_| |_(_)
                                           
                                           

// 		`)
// 		fmt.Println("fin du jeu vous avez gagner le mot était bien", hangman.Mot)
// 	hangman.IsRunning = false
	
// }

// func (hangman *HANGMAN) gameOver() {
// 	fmt.Println(` 
// 		   ____                          ___                 
// 		  / ___| __ _ _ __ ___   ___    / _ \__   _____ _ __ 
// 		 | |  _ / _' | '_ ' _ \ / _ \  | | | \ \ / / _ \ '__|
// 		 | |_| | (_| | | | | | |  __/  | |_| |\ V /  __/ |
// 	    	  \____|\__,_|_| |_| |_|\___|   \___/  \_/ \___|_|
                                                                                              

// 	`)

// 	fmt.Println("fin du jeu le mot était : ", hangman.Mot)
// 	hangman.IsRunning = false
	
// }