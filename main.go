package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/twitter", twitter)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		renderTemplate(w, "index", r)
		return
	}
	http.Error(w, "404 Not Found", http.StatusNotFound)
}

func twitter(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "twitter", r)
}

func renderTemplate(w http.ResponseWriter, fname string, r *http.Request) {
	t, _ := template.ParseFiles("templates/base.html", "templates/"+fname+".html")
	t.ExecuteTemplate(w, "base", "")
}
