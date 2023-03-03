package controller

import (
	"book/pb"
	"book/pb/book"
	"context"
)

func (b *BookController) Update(ctx context.Context, req *book.BookUpdateRequest) (res *pb.OperationResponse, err error) {
	res, err = b.service.Update(ctx, req)
	return
}
