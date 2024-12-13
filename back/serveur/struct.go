package hangmanweb


type Donnees struct {   // structure contenant les donnée a afficher dans les templates html
	Essai      int      //nombre d'essai entre 0 et 10 inclu 
	Mot        string   // mot a deviner DEBBUGAGE A SUPPRIMER 
	MotInconnu string   // Mot inconnue donc afficher sous la forme "a _ _ a _ _ " qui correspond a la variable MotInconnu de type []string
	Usedletter string   // lettre utiliser 
}

type HANGMANWEB struct {
	//variable contenant les chemins d'accés au tamplates
	TemplateHome         string
	TemplatePlay         string
	TemplateWin          string
	TemplateLoose        string

	//variables pour les list de mot easy, medium et hard 
	WordlsitEasy       []string // liste de mot easy
	WordlistMedium     []string  // liste de mot medium
	wordlistHard       []string  // liste de mot hard

	// mot a deviner pour chaque niveau type string
	MotEasy            string
	MotMedium          string
	MotHard            string
	Mot                string

	// tableau de string pour comparer les lettre du mot a deviner et celle entrée par le joueur 
	MotAdeviner        []string   // mot a deviener de type [c a n a r d]
	MotIconnu          []string    // mot inconnu de type  [c a _ a _ d]

	
	// variable des requêtes lettre entrée et  niveau choisis et used letter
	LevelActual        string
	Erreur             int
	Usedletter         string




}
