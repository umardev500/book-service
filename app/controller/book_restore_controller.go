package controller

import (
	"book/pb"
	"book/pb/book"
	"context"
)

func (b *BookController) Restore(ctx context.Context, req *book.BookRestoreRequest) (res *pb.OperationResponse, err error) {
	res, err = b.service.Restore(ctx, req)
	return
}
