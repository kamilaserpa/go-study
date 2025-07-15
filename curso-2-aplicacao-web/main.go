package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", handlerIndex)
	http.ListenAndServe(":8000", nil)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul", Preco: 30, Quantidade: 5},
		{Nome: "Notebook", Descricao: "Dell 14\" ", Preco: 2999, Quantidade: 4},
		{"Tênis", "Confortável", 149, 3},
		{"Fone de ouvido", "Com fio", 59, 2},
	}
	templates.ExecuteTemplate(w, "Index", produtos)
}
