package main

import (
	"log"
	"net/http"

	"workshop/internal/handler"

	"github.com/go-chi/chi/v5"
)

func main() {
	h := handler.NewHandler()
	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Print("Starting server")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)

	log.Print("Shutting server down")
}
