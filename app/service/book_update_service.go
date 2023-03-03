package service

import (
	"book/pb"
	"book/pb/book"
	"context"
)

func (b *bookService) Update(ctx context.Context, req *book.BookUpdateRequest) (res *pb.OperationResponse, err error) {
	res, err = b.repo.Update(ctx, req)
	return
}
