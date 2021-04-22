package repositories

import (
	models "webserver/app/Models"

	"gorm.io/gorm"
)


type BookRepository struct {
	SqlConn *gorm.DB
}

func (repository *BookRepository) Create(paylod *models.Book) {
	repository.SqlConn.Create(&paylod)
}
