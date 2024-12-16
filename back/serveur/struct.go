package hangmanweb


type DataStruct struct {   // structure contenant les donnée a afficher dans les templates html
	Attemps      int      //nombre d'essai entre 0 et 10 inclu 
	Word        string   // Word a deviner DEBBUGAGE A SUPPRIMER 
	HiddenWord string   // Word inconnue donc afficher sous la forme "a _ _ a _ _ " qui correspond a la variable WordInconnu de type []string
	Usedletter string   // lettre utiliser 
}

type HANGMANWEB struct {
	//variable contenant les chemins d'accés au tamplates
	TemplateHome         string
	TemplatePlay         string
	TemplateWin          string
	TemplateLoose        string

	//variables pour les list de Word easy, medium et hard 
	WordlsitEasy       []string // liste de Word easy
	WordlistMedium     []string  // liste de Word medium
	wordlistHard       []string  // liste de Word hard

	// Word a deviner pour chaque niveau type string
	WordEasy            string
	WordMedium          string
	WordHard            string
	Word                string

	// tableau de string pour comparer les lettre du Word a deviner et celle entrée par le joueur 
	SecretWord        []string   // Word a deviener de type [c a n a r d]
	HiddenWord          []string    // Word inconnu de type  [c a _ a _ d]

	
	// variable des requêtes lettre entrée et  niveau choisis et used letter
	LevelActual        string
	Error             int
	Usedletter         string




}
