package hangmanweb

import (
	"html/template"
	"net/http"
)









func (h *HANGMANWEB) Home(w http.ResponseWriter, r *http.Request) {
	h.RequestAndTemplate(w, r, 0)
	
}

func (h *HANGMANWEB) Play(w http.ResponseWriter, r *http.Request) {
	h.RequestAndTemplate(w, r, 1)
}
	

func (h *HANGMANWEB) Win(w http.ResponseWriter, r *http.Request) {
	h.RequestAndTemplate(w, r, 2)
}

func (h *HANGMANWEB) Loose(w http.ResponseWriter, r *http.Request) {
	h.RequestAndTemplate(w, r, 3)
}



func(h *HANGMANWEB) RequestAndTemplate(w http.ResponseWriter, r *http.Request, TmplNb int)  {
	r.ParseForm()
	tmpl := template.Must(template.ParseFiles(h.TmplList[TmplNb]))
	r.Form.Get("lettre")
	r.Form.Get("level")
	tmpl.Execute(w, Data)
	
}