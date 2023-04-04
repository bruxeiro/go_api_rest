package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func main() {

	fmt.Println("Iniciando projeto Go_Api_REST")
}
