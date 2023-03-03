package controller

import (
	"book/domain"
	"book/pb/book"
)

type BookController struct {
	service domain.BookService
	book.UnimplementedBookServiceServer
}

func NewBookController(service domain.BookService) *BookController {
	return &BookController{
		service: service,
	}
}
