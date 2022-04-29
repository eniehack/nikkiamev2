package handler

import (
	"github.com/eniehack/nikkiamev2/internal/config"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)


type Handler struct {
	DB *gorm.DB
	Validator *validator.Validate
	SessionStore *sessions.CookieStore
	Config *config.Config
}
