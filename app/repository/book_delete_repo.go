package repository

import (
	"book/pb"
	"book/pb/book"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (b *bookRepo) Delete(ctx context.Context, req *book.BookDeleteRequest) (res *pb.OperationResponse, err error) {
	isAffected := false
	filter := bson.M{"book_id": req.BookId}
	resp, err := b.book.DeleteOne(ctx, filter)
	if err != nil {
		return
	}

	if resp.DeletedCount > 0 {
		isAffected = true
	}

	res = &pb.OperationResponse{IsAffected: isAffected}

	return
}
