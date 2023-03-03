package controller

import (
	"book/pb/book"
	"context"
)

func (b *BookController) GetBooks(ctx context.Context, req *book.BookFindAllRequest) (res *book.BookFindAllResponse, err error) {
	res, err = b.service.GetBooks(ctx, req)
	return
}
