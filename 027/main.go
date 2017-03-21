package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/flowers/", flowers)
	http.HandleFunc("/flowers.jpg", chien)
	http.ListenAndServe(":9000", nil)


}

func foo(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "flowers smell beautiful")
}

func flowers(w http.ResponseWriter, req *http.Request){
	tpl, err := template.ParseFiles("flowers.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "flowers.gohtml", nil)
}

func chien(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "flowers.jpg")
}