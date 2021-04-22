package main

import (
	"sync"
	databases "webserver/app/Config"
	controllers "webserver/app/Controller"
	repositories "webserver/app/Repositories"
)

type IServiceContainer interface {
	InjectBookController() controllers.BookController
}

type kernel struct{}

func (k *kernel) InjectBookController() controllers.BookController {

	SqlConn := databases.Connect()

	playerRepository := &repositories.BookRepository{SqlConn: SqlConn}
	playerController := controllers.BookController{Repository: playerRepository}

	return playerController
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
