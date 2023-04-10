package routes

import (
	"log"
	"net/http"

	"go_api_rest/controllers"

	"github.com/gorilla/mux"
)

func HandleResquest() {

	//Criando router para o gorilla mux
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.CriaPerso).Methods("Post")
	r.HandleFunc("/api/personalidades", controllers.TodasPerso).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPerso).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPerso).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPerso).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))

}
