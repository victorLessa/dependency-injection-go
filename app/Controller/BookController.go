package controllers

import (
	"encoding/json"
	"net/http"

	interfaces "webserver/app/Interfaces"
	models "webserver/app/Models"
	repositories "webserver/app/Repositories"
)

type BookController struct {
	interfaces.ICommonController
	Repository *repositories.BookRepository
}

// Cria livros
func (c *BookController) Create(w http.ResponseWriter, r *http.Request) {
	books := models.Book{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&books); err != nil {
		RespondJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	
	c.Repository.Create(&books)
	RespondJSON(w, http.StatusOK, books) 
}