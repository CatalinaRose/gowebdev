package main

import (
	"net/http"
	"html/template"
)




type person struct {
	First string
	Last string
	Saying string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":9003", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Gata's",
		Last: "Pot Head's",
		Saying: "Smoke Weed everday!",
	}

	p2 := person{
		First: "Miss",
		Last: "G",
		Saying: "Welcome to Good Burger, can I take your order?",
	}

	xp := []person{p1, p2}
	tpl.ExecuteTemplate(w, "index.gohtml", xp)
}

