package main

import (
	"log"
	"net/http"

	"flag"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/eniehack/nikkiamev2/internal/config"
)

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "./config.toml", "please specify config file's place")
	flag.Parse()

	configData, err := config.LoadConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config file: %s", err)
	}

	db, err := config.SetupDatabase(configData)
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	http.ListenAndServe(":3000", r)
}

