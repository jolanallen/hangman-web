package hangman

import "fmt"


func (hangman *HANGMAN) Hangman() {
  
  var etat0 = hangman.TabHang[0]
  var etat1 = hangman.TabHang[1]
  var etat2 = hangman.TabHang[2]
  var etat3 = hangman.TabHang[3]
  var etat4 = hangman.TabHang[4]
  var etat5 = hangman.TabHang[5]
  var etat6 = hangman.TabHang[6]
  var etat7 = hangman.TabHang[7]
  var etat8 = hangman.TabHang[8]
  var etat9 = hangman.TabHang[9]
 
  if hangman.Erreur == 0 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat0, "                       Attempts :", hangman.Erreur, "/9")
  }
  if hangman.Erreur == 1 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat1, "                       Attempts :", hangman.Erreur, "/9")
  }
  if hangman.Erreur == 2 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat2, "                       Attempts :", hangman.Erreur, "/9")
  }
  if hangman.Erreur == 3 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat3, "                       Attempts :", hangman.Erreur, "/9")
  }
  if hangman.Erreur == 4 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat4, "                       Attempts :", hangman.Erreur, "/9")
  }
  if hangman.Erreur == 5 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat5, "                       Attempts :", hangman.Erreur, "/9")
  }
  if hangman.Erreur == 6 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat6, "                       Attempts :", hangman.Erreur, "/9")
  }
  if hangman.Erreur == 7 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat7, "                       Attempts :", hangman.Erreur, "/9")
  }
  if hangman.Erreur == 8 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat8, "                       Attempts :", hangman.Erreur, "/9")
  }
  if hangman.Erreur == 9 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat9, "                       Attempts :", hangman.Erreur, "/9")
  }
}
