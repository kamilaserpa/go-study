package controllers

import (
	"alura/curso-2-aplicacao-web/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	// Linka com tag definida no html {{define "Index"}}
	templates.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func HandlerProdutoNovo(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "ProdutoNovo", nil)
}

func InsertProdutoDb(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço", err)
			http.Error(w, "Erro na conversão do preço: "+err.Error(), http.StatusBadRequest)
			return
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err)
			return
		}

		models.CriarProdutoDb(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func DeleteProdutoDb(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletarProdutoDb(idDoProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func HandlerEditProduto(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.BuscaProdutoDB(idDoProduto)
	templates.ExecuteTemplate(w, "EditProduto", produto)
}

func UpdateProdutoDb(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		// Converte string para
		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID: ", err)
			return
		}

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preco: ", err)
			return
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err)
			return
		}
		models.AtualizarProdutoDb(idConvertido, nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
