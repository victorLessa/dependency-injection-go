package interfaces

import models "webserver/app/Models"

type IBookRepository interface {
	Create(paylod models.Book)
}
