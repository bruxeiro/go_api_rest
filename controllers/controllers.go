package controllers

import (
	"encoding/json"
	"fmt"
	"go_api_rest/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPerso(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

// Função para retornar personalidades usando o gorila mux
func RetornaPerso(w http.ResponseWriter, r *http.Request) {
	//Passando a request "r" em parametros para Vars
	vars := mux.Vars(r)
	//Puxando apenas o id contido em vars
	id := vars["id"]

	//Comparar 1 personalidade por vez baseada no ID
	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id {
			json.NewEncoder(w).Encode(personalidade)
		}
	}

}
