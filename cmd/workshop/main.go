package main

import (
	"log"
	"net/http"

	"workshop/internal/api/jokes"
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

	apiClient := jokes.NewJokeClient(cfg.JokeURL)

	h := handler.NewHandler(apiClient)

	r := chi.NewRouter()

	r.Get("/joke", h.Joke)
	path := cfg.Host + ":" + cfg.Port

	log.Print("Starting server")
	err = http.ListenAndServe(path, r)
	log.Fatal(err)

	log.Print("Shutting server down")
}
