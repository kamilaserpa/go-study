package main

import (
	"fmt"

	"github.com/kamilaserpa/go-study/models"
	"github.com/kamilaserpa/go-study/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{
			Id:       1,
			Nome:     "Albert Einstein",
			Historia: "Físico teórico que desenvolveu a teoria da relatividade.",
		},
		{
			Id:       2,
			Nome:     "Marie Curie",
			Historia: "Pioneira na pesquisa sobre radioatividade.",
		},
	}

	fmt.Println("Iniciando o servidor Rest com Go!")
	routes.HandleRequest()
}
