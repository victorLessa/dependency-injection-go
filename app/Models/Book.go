package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Author string `json:"author"`
	Age    int    `json:"age"`
}