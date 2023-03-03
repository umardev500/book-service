package repository

import (
	"book/helper"
	"book/pb"
	"book/pb/book"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (b *bookRepo) Update(ctx context.Context, req *book.BookUpdateRequest) (res *pb.OperationResponse, err error) {
	isAffected := false
	filter := bson.M{"book_id": req.BookId}

	uploaderPayload := bson.D{}
	if req.Payload.Uploader != nil {
		uploaderPayload = bson.D{
			{Key: "uploader.user_id", Value: req.Payload.Uploader.UserId},
			{Key: "uploader.user", Value: req.Payload.Uploader.User},
		}

		helper.NoEmpty(uploaderPayload, &uploaderPayload)
	}

	bookPayload := bson.D{
		{Key: "title", Value: req.Payload.Title},
		{Key: "author", Value: req.Payload.Author},
		{Key: "publisher", Value: req.Payload.Publisher},
		{Key: "pages_total", Value: req.Payload.PagesTotal},
		{Key: "qty", Value: req.Payload.Qty},
		{Key: "cover", Value: req.Payload.Cover},
		{Key: "description", Value: req.Payload.Description},
		{Key: "location_id", Value: req.Payload.LocationId},
	}

	helper.NoEmpty(bookPayload, &bookPayload)

	// editor
	editorPayload := bson.M{}
	editors := req.Payload.Editors
	if len(editors) > 0 {
		editor := editors[0]
		editorPayload = bson.M{
			"editors": bson.M{
				"user_id": editor.UserId,
				"user":    editor.User,
			},
		}
	}

	payload := bson.D{}
	payload = append(payload, bookPayload...)
	payload = append(payload, uploaderPayload...)
	payload = append(payload, bson.E{Key: "updated_at", Value: helper.GetTime()})

	// set := bson.M{"$set": editorPaylod}
	opts := options.Update().SetUpsert(false)
	resp, err := b.book.UpdateOne(ctx, filter, bson.M{"$set": payload, "$addToSet": editorPayload}, opts)
	if resp.ModifiedCount > 0 {
		isAffected = true
	}

	res = &pb.OperationResponse{IsAffected: isAffected}

	return
}
