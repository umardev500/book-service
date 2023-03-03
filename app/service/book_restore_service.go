package service

import (
	"book/pb"
	"book/pb/book"
	"context"
)

func (b *bookService) Restore(ctx context.Context, req *book.BookRestoreRequest) (res *pb.OperationResponse, err error) {
	res, err = b.repo.Restore(ctx, req)
	return
}
