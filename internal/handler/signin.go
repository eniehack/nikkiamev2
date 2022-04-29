package handler

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/eniehack/nikkiamev2/internal/model"
)

type signinParam struct {
	userID string `validate:"required,user_id"`
	password string `validate:"required"`
}

func (h *Handler) SigninIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/signin.html.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Signin(w http.ResponseWriter, r *http.Request) {
	var user model.Users
	session, _ := h.SessionStore.Get(r, "nikkiame-session")

	log.Printf("%v", session.Values["user_id"])
	userid, _ := session.Values["user_id"].(string)
	if userid != "" {
		http.Redirect(w, r, "/", 302)
	}

	r.ParseForm()
	params := signinParam{
		userID: r.FormValue("user_id"),
		password: r.FormValue("password"),
	}
	h.Validator.Struct(params)
	h.DB.Where(&model.Users{UserID: params.userID}).First(&user)

	if match, err := argon2id.ComparePasswordAndHash(params.password, user.Password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("signin/argon2id: %v", err)
		return
	} else if match != true {
		tmpl, err := template.ParseFiles("templates/signin.html.tmpl")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		if err := tmpl.Execute(w, struct{UserID string}{UserID: params.userID}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	user.UpdatedAt = time.Now()
	h.DB.Save(&user)

	session.Values["user_id"] = user.Ulid

	if err := session.Save(r,w); err != nil {
		log.Fatalf("signin/session save: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", 200)
	return
}
