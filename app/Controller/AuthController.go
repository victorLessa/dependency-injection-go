package controllers

import (
	"encoding/json"
	"net/http"
	helpers "webserver/app/Helpers"
	interfaces "webserver/app/Interfaces"
	models "webserver/app/Models"
)


type AuthController struct {
	interfaces.ICommonController
}

func (a *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	var c map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		RespondJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	token, _ := helpers.TokenAuth(c)
	response := models.Token{}
	response.Token = token
	
	RespondJSON(w, http.StatusOK, response)
}