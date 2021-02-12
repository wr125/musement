package handler

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.html"))
}

//Index export to main.go
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tpl.ExecuteTemplate(w, "index.html", nil)

}
func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	Name := r.FormValue("City_Name")
	tpl.ExecuteTemplate(w, "result.html", Name)

}
