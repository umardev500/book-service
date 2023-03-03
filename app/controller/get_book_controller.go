package controller

import (
	"book/pb/book"
	"context"
)

func (b *BookController) GetBook(ctx context.Context, req *book.BookFindOneRequest) (res *book.BookFindOneResponse, err error) {
	res, err = b.service.GetBook(ctx, req)

	return
}
