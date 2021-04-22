package migration

import (
	database "webserver/app/Config"
	model "webserver/app/Models"
)

func RunMigrations() {
	db := database.Connect()
	db.AutoMigrate(&model.Book{})
}