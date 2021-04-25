package main

import (
	"fmt"
	"net/http"
	"sync"
	helpers "webserver/app/Helpers"

	"github.com/go-chi/chi"
	jwtauth "github.com/go-chi/jwtauth/v5"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	var bookController = ServiceContainer().InjectBookController()
	var authController = ServiceContainer().InjectAuthController()

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		_, jwtToken := helpers.TokenAuth(map[string]interface{}{})
		r.Use(jwtauth.Verifier(jwtToken))

		r.Use(jwtauth.Authenticator)
		r.Post("/books", bookController.Create)
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	})

	r.HandleFunc("/login", authController.SignIn)

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}