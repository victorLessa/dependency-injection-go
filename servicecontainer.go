package main

import (
	"sync"
	databases "webserver/app/Config"
	controllers "webserver/app/Controller"
	repositories "webserver/app/Repositories"
)

type IServiceContainer interface {
	InjectBookController() controllers.BookController
	InjectAuthController() controllers.AuthController
}

type kernel struct{}

func (k *kernel) InjectBookController() controllers.BookController {

	SqlConn := databases.Connect()

	bookRepository := &repositories.BookRepository{SqlConn: SqlConn}
	bookController := controllers.BookController{Repository: bookRepository}

	return bookController
}
func (k *kernel) InjectAuthController() controllers.AuthController {
	userController := controllers.AuthController{}

	return userController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
