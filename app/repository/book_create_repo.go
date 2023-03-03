package repository

import (
	"book/pb/book"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (b *bookRepo) Create(ctx context.Context, req *book.BookCreateRequest) (res *book.BookCreateResponse, err error) {

	bookPayload := bson.D{
		{Key: "book_id", Value: req.Title},
		{Key: "title", Value: req.Title},
		{Key: "author", Value: req.Author},
		{Key: "publisher", Value: req.Publisher},
		{Key: "pages_total", Value: req.PagesTotal},
		{Key: "qty", Value: req.Qty},
		{Key: "cover", Value: req.Cover},
		{Key: "description", Value: req.Description},
		{Key: "location_id", Value: req.LocationId},
	}

	_, err = b.book.InsertOne(ctx, bookPayload)
	// fmt.Println("error", err)

	return
}
