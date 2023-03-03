package service

import (
	"book/pb/book"
	"context"
)

func (b *bookService) GetBook(ctx context.Context, req *book.BookFindOneRequest) (res *book.BookFindOneResponse, err error) {
	res, err = b.repo.GetBook(ctx, req)
	return
}
