package repository

import (
	"book/pb"
	"book/pb/book"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (b *bookRepo) Restore(ctx context.Context, req *book.BookRestoreRequest) (res *pb.OperationResponse, err error) {
	isAffected := false
	filter := bson.M{"book_id": req.BookId}
	set := bson.M{"$unset": bson.M{"deleted_at": nil}}
	resp, err := b.book.UpdateOne(ctx, filter, set)
	if resp.ModifiedCount > 0 {
		isAffected = true
	}

	res = &pb.OperationResponse{IsAffected: isAffected}

	return
}
