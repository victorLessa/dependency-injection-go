package interfaces

import "net/http"

type ICommonController interface {
	RespondJSON(w http.ResponseWriter, status int, payload interface{})
	RespondError(w http.ResponseWriter, code int, message string)
}