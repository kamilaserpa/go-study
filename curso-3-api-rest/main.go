package main

import (
	"fmt"

	"github.com/kamilaserpa/go-study/database"
	"github.com/kamilaserpa/go-study/routes"
)

func main() {
	database.ConnectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go!")
	routes.HandleRequest()
}
