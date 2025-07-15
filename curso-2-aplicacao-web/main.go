package main

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Index", nil)
}
