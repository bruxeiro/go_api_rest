package controllers

import (
	"encoding/json"
	"fmt"
	"go_api_rest/database"
	"go_api_rest/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPerso(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

// Função para retornar personalidades usando o gorila mux
func RetornaPerso(w http.ResponseWriter, r *http.Request) {
	//Passando a request "r" em parametros para Vars
	vars := mux.Vars(r)
	//Puxando apenas o id contido em vars
	id := vars["id"]

	var p models.Personalidade
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}
