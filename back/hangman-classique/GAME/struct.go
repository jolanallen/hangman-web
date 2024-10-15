package hangman

import "strings"



type HANGMAN struct {          // vatiable utiliser dans la structure du jeu
	IsRunning     bool
	wordIsGood    bool
	TabMots       []string
	Mot           string
	MotAdeviner   []string
	TabHang       []string
	etat          strings.Builder
	motIconnu     []string
	randomNb      int
	erreur       int
	lettreIsGood bool
	wordFile     string
	lettre       string
	UsedLetter   []string
	medium       bool
	hard         bool
}

