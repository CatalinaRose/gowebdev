package main


import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	m := map[string]bool{}
	m["Pot"] = true
	m["Babe"] = true
	m["Gata"] = true
	tpl.ExecuteTemplate(w, "index.gohtml", nil)

}

