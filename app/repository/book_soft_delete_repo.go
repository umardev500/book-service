package repository

import (
	"book/helper"
	"book/pb"
	"book/pb/book"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (b *bookRepo) SoftDelete(ctx context.Context, req *book.BookDeleteRequest) (res *pb.OperationResponse, err error) {
	isAffected := false
	filter := bson.M{"book_id": req.BookId, "deleted_at": bson.M{"$eq": nil}}
	payload := bson.M{"deleted_at": helper.GetTime()}
	set := bson.M{"$set": payload}
	resp, err := b.book.UpdateOne(ctx, filter, set)
	if err != nil {
		return
	}

	if resp.ModifiedCount > 0 {
		isAffected = true
	}

	res = &pb.OperationResponse{IsAffected: isAffected}

	return
}
