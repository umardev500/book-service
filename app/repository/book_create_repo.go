package repository

import (
	"book/helper"
	"book/pb/book"
	"book/variable"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (b *bookRepo) Create(ctx context.Context, req *book.BookCreateRequest) (res *book.BookCreateResponse, err error) {
	uploaderPayload := bson.D{}
	if req.Uploader != nil {
		uploaderPayload = bson.D{
			{Key: "user_id", Value: req.Uploader.UserId},
			{Key: "user", Value: req.Uploader.User},
		}

		helper.NoEmpty(uploaderPayload, &uploaderPayload)
	}

	bookPayload := bson.D{
		{Key: "book_id", Value: helper.GetTime(&variable.UnixNano)},
		{Key: "title", Value: req.Title},
		{Key: "author", Value: req.Author},
		{Key: "publisher", Value: req.Publisher},
		{Key: "pages_total", Value: req.PagesTotal},
		{Key: "qty", Value: req.Qty},
		{Key: "cover", Value: req.Cover},
		{Key: "description", Value: req.Description},
		{Key: "location_id", Value: req.LocationId},
		{Key: "uploader", Value: uploaderPayload},
	}

	helper.NoEmpty(bookPayload, &bookPayload)

	payload := bson.D{}
	payload = append(payload, bookPayload...)
	payload = append(payload, uploaderPayload...)
	payload = append(payload, bson.E{Key: "created_at", Value: helper.GetTime(nil)})

	res = &book.BookCreateResponse{}
	_, err = b.book.InsertOne(ctx, payload)

	return
}
