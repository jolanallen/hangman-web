package hangman

import "fmt"


func (hangman *HANGMAN) hangman() {
  
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
 
  if hangman.erreur == 0 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat0, "                       Attempts :", hangman.erreur, "/9")
  }
  if hangman.erreur == 1 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat1, "                       Attempts :", hangman.erreur, "/9")
  }
  if hangman.erreur == 2 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat2, "                       Attempts :", hangman.erreur, "/9")
  }
  if hangman.erreur == 3 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat3, "                       Attempts :", hangman.erreur, "/9")
  }
  if hangman.erreur == 4 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat4, "                       Attempts :", hangman.erreur, "/9")
  }
  if hangman.erreur == 5 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat5, "                       Attempts :", hangman.erreur, "/9")
  }
  if hangman.erreur == 6 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat6, "                       Attempts :", hangman.erreur, "/9")
  }
  if hangman.erreur == 7 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat7, "                       Attempts :", hangman.erreur, "/9")
  }
  if hangman.erreur == 8 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat8, "                       Attempts :", hangman.erreur, "/9")
  }
  if hangman.erreur == 9 {
    fmt.Println("                        Lettres déja utilisées :", hangman.UsedLetter)
    fmt.Println(etat9, "                       Attempts :", hangman.erreur, "/9")
  }
}
