package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	models "webserver/app/Models"

	"github.com/go-chi/chi"
	jwtauth "github.com/go-chi/jwtauth/v5"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func TokenAuth(paylod map[string]interface{}) (string, *jwtauth.JWTAuth) {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
		_, tokenString, _ := tokenAuth.Encode(paylod)
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
	return tokenString, tokenAuth
}

func (router *router) InitRouter() *chi.Mux {

	bookController := ServiceContainer().InjectBookController()

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		_, jwtToken := TokenAuth(map[string]interface{}{})
		r.Use(jwtauth.Verifier(jwtToken))

		r.Use(jwtauth.Authenticator)
		r.Post("/books", bookController.Create)
	})

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var c map[string]interface{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&c); err != nil {
			bookController.RespondJSON(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		token, _ := TokenAuth(c)
		response := models.Token{}
		response.Token = token
		
		bookController.RespondJSON(w, http.StatusOK, response)
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	})

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