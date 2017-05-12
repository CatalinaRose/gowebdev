package main

import (
	"net/http"
	"html/templates"
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
	http.ListenAndServe(":5058", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Gata",
		Last: "Arroyo",
		Saying: "Need to say, stfu will u lol",

	}

	p2 := person{
		First: "Cat",
		Last: "maggot",
		Saying: "Show me the money!",
	}

	xp := []person{p1, p2}
	tpl.ExecuteTemplate(w,"index.gohtml", xp)
}

