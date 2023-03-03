package controller

import (
	"book/pb"
	"book/pb/book"
	"context"
)

func (b *BookController) SoftDelete(ctx context.Context, req *book.BookDeleteRequest) (res *pb.OperationResponse, err error) {
	res, err = b.service.SoftDelete(ctx, req)
	return
}
