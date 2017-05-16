package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":9002", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name: "Yo, Yo",
		Value: "tecate",
		Path: "/",
		HttpOnly: true,
	}
	http.SetCookie(w, c)
	io.WriteString(w, "<h1>hello, this is my favorite beer!</h1>")
}

func bar(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("wasssup")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	io.WriteString(w, c.Value)
}
