package controller

import (
	"book/pb"
	"book/pb/book"
	"context"
)

func (b *BookController) Delete(ctx context.Context, req *book.BookDeleteRequest) (res *pb.OperationResponse, err error) {
	res, err = b.service.Delete(ctx, req)
	return
}
