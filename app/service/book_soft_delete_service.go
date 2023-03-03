package service

import (
	"book/pb"
	"book/pb/book"
	"context"
)

func (b *bookService) SoftDelete(ctx context.Context, req *book.BookDeleteRequest) (res *pb.OperationResponse, err error) {
	res, err = b.repo.SoftDelete(ctx, req)

	return
}
