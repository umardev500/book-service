package controller

import (
	"book/domain"
	"book/pb/book"
	"context"
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

func (b *BookController) GetBooks(ctx context.Context, req *book.BookFindAllRequest) (res *book.BookFindAllResponse, err error) {
	res, err = b.service.GetBooks(ctx, req)
	return
}
