package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func contact(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func about(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}



