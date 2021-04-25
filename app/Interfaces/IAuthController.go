package interfaces

import "net/http"

type IAuthController interface {
	SingIn(w http.ResponseWriter, r *http.Request)
}