package routes

import (
	"log"
	"net/http"

	"github.com/kamilaserpa/go-study/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil)) // Qualquer problema na inicialização será logado
}
