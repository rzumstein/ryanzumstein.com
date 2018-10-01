package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/twitter", twitter)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile("templates/index.html")
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, f)
}

func twitter(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile("templates/twitter.html")
	t, _ := template.ParseFiles("templates/twitter.html")
	t.Execute(w, f)
}
