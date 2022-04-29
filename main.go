package main

import (
	"log"
	"net/http"

	"flag"

	"github.com/eniehack/nikkiamev2/internal/config"
	"github.com/eniehack/nikkiamev2/internal/handler"
	"github.com/eniehack/nikkiamev2/internal/validation"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
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

	v := validator.New()
	v.RegisterValidation("user_id", validation.UserIDValidator)

	store := sessions.NewCookieStore([]byte(configData.Session.Secret))

	h := &handler.Handler{
		DB: db,
		Validator: v,
		SessionStore: store,
		Config: configData,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	r.Get("/signin", h.SigninIndex)
	r.Post("/signin", h.Signin)
	r.Get("/@{userid}/feed.xml", h.UserFeed)

	http.ListenAndServe(":3000", r)
}
