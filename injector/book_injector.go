package injector

import (
	"book/app/controller"
	"book/app/repository"
	"book/app/service"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewBookInjector(book *mongo.Collection) *controller.BookController {
	bookRepo := repository.NewBookRepo(book)
	bookService := service.NewBookService(bookRepo)
	return controller.NewBookController(bookService)
}
