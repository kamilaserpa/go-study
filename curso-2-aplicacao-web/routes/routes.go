package routes

import (
	"alura/curso-2-aplicacao-web/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.HandlerIndex)
	http.HandleFunc("/novo-produto", controllers.HandlerProdutoNovo)
	http.HandleFunc("/insert-produto-db", controllers.InsertDb)
}
