package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", dogs)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":9095", nil)
}

func dogs(w http.ResponseWriter, req *http.Request) {
	 tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
