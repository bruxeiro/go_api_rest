package main

import (
	"fmt"
	"go_api_rest/models"
	"go_api_rest/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando projeto Go_Api_REST")
	routes.HandleResquest()
}
