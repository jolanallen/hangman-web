package hangman

import (
    

    
)
 

func (hangman *HANGMAN) Run() {
    hangman.wordIsGood = false
    for hangman.IsRunning {
        hangman.Hangman()
        hangman.testword()
        
    }
}



func (hangman *HANGMAN) TestLetter(lettre string) {
    hangman.lettreIsGood = false
    
    for l := range hangman.MotAdeviner {
        if lettre == string(hangman.MotAdeviner[l]) {
            hangman.lettreIsGood = true
            hangman.MotIconnu[l] = hangman.MotAdeviner[l] // Utilise la lettre originale pour l'affichage
        }
    }
    if !hangman.lettreIsGood {
        hangman.Erreur++
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