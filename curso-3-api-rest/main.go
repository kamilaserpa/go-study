package main

import (
	"fmt"

	"github.com/kamilaserpa/go-study/models"
	"github.com/kamilaserpa/go-study/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{
			ID:        1,
			Nome:      "Albert Einstein",
			Historia:  "Físico teórico que desenvolveu a teoria da relatividade.",
			Descricao: "Considerado um dos cientistas mais influentes do século XX.",
		},
		{
			ID:        2,
			Nome:      "Marie Curie",
			Historia:  "Pioneira na pesquisa sobre radioatividade.",
			Descricao: "Primeira mulher a ganhar um Prêmio Nobel e a única pessoa a ganhar em duas áreas científicas diferentes.",
		},
	}

	fmt.Println("Iniciando o servidor Rest com Go!")
	routes.HandleRequest()
}
