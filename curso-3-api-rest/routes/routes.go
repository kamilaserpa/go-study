package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kamilaserpa/go-study/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalidadePorId).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r)) // Qualquer problema na inicialização será logado
}
