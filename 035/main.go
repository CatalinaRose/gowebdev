package main

import (
	"net/http"
	"io"
)


func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8041", nil)
}
func index (w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name: "your a fag",
		Value: "What did u say?",
		Path: "/",
		HttpOnly: true,
	}
	http.SetCookie(w, c)
	io.WriteString(w, "<h1>Yo!</h1>")

}