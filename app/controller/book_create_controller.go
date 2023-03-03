package controller

import (
	"book/pb/book"
	"context"
)

func (b *BookController) Create(ctx context.Context, req *book.BookCreateRequest) (res *book.BookCreateResponse, err error) {
	res, err = b.service.Create(ctx, req)
	return
}
