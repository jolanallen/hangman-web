package hangmanweb





type Donnees struct {
	Essai   int
	Mot     string
	MotInconnu string
	Usedletter string
	
	
}



type HANGMANWEB struct {
	TemplateHome  string
	TemplatePlay  string
	TemplateWin   string 
	TemplateLoose string
	TmplNb        int
	TmplList      []string
	Lettre      string
	Level        string

}
