package service

import (
	"book/pb/book"
	"context"
)

func (b *bookService) GetBooks(ctx context.Context, req *book.BookFindAllRequest) (res *book.BookFindAllResponse, err error) {
	res, err = b.repo.GetBooks(ctx, req)
	return
}
