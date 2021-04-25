package helpers

import (
	"fmt"

	"github.com/go-chi/jwtauth/v5"
)

func TokenAuth(paylod map[string]interface{}) (string, *jwtauth.JWTAuth) {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, _ := tokenAuth.Encode(paylod)
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
	return tokenString, tokenAuth
}