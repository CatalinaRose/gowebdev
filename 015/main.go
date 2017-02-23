package main

import(
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", link)
	http.ListenAndServe(":8083", nil)
}

func link(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "link.gohtml", "Im am not of this world")
}

