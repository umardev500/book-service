package service

import (
	"book/pb"
	"book/pb/book"
	"context"
)

func (b *bookService) Delete(ctx context.Context, req *book.BookDeleteRequest) (res *pb.OperationResponse, err error) {
	res, err = b.repo.Delete(ctx, req)

	return
}
