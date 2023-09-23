package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prranavv/go-odin-2/routes"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", routes.Mainpage)
	r.Post("/new-message", routes.NewMessage)
	http.ListenAndServe(":8080", r)
}
