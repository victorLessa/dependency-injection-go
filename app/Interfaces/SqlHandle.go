package interfaces

import "gorm.io/gorm"

type SqlHandler interface {
	Db(statement string) *gorm.DB
}