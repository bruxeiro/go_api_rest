package routes

import (
	"log"
	"net/http"

	"go_api_rest/controllers"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPerso)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
