package repository

import (
	"book/domain"
	"book/pb/book"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (b *bookRepo) GetBook(ctx context.Context, req *book.BookFindOneRequest) (res *book.BookFindOneResponse, err error) {
	var item domain.Book

	filter := bson.M{"book_id": req.BookId}
	err = b.book.FindOne(ctx, filter).Decode(&item)
	if err != nil && err != mongo.ErrNoDocuments {
		return
	}

	if err == mongo.ErrNoDocuments {
		res = &book.BookFindOneResponse{IsEmpty: true}
		err = nil
		return
	}

	bookItem := b.parseBook(item)
	res = &book.BookFindOneResponse{
		Payload: bookItem,
	}

	return
}
