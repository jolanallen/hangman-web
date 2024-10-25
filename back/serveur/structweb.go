package hangmanweb


import (
	"net/http"
)





type HANGMANWEB struct {
	CheminTemplate string
	Donnees        interface{}
	rw             http.ResponseWriter
	r             *http.Request
	cheminTemplate string
	
	

}
