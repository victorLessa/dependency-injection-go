package controllers

import (
	"encoding/json"
	"net/http"

	models "webserver/app/Models"
	repositories "webserver/app/Repositories"
)

type BookController struct {
	Repository *repositories.BookRepository
}

// Cria livros
func (c *BookController) Create(w http.ResponseWriter, r *http.Request) {
	books := models.Book{}

		decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&books); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	
	c.Repository.Create(&books)
	respondJSON(w, http.StatusOK, books) 
}