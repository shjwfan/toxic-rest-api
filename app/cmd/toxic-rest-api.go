package main

import (
	"net/http"

	"github.com/anotherandrey/token-rest-api/internal/app/api"
)

func main() {
	api := api.NewApi()

	http.ListenAndServe(":8080", api.Router)
	// http.ListenAndServeTLS(":8080", "", api.Router)
}
