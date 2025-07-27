package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kamilaserpa/go-study/controllers"
	"github.com/kamilaserpa/go-study/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalidadePorId).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CriaNovaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r)) // Qualquer problema na inicialização será logado
}
