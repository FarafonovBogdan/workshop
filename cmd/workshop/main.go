package main

import (
	"log"
	"net/http"

	"workshop/internal/config"
	"workshop/internal/handler"

	"github.com/go-chi/chi/v5"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {

	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	h := handler.NewHandler()

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)
	path := cfg.Host + ":" + cfg.Port

	log.Print("Starting server")
	err = http.ListenAndServe(path, r)
	log.Fatal(err)

	log.Print("Shutting server down")
}
