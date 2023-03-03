package service

import (
	"book/pb/book"
	"context"
)

func (b *bookService) Create(ctx context.Context, req *book.BookCreateRequest) (res *book.BookCreateResponse, err error) {
	res, err = b.repo.Create(ctx, req)

	return
}
