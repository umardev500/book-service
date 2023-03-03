package repository

import (
	"book/domain"
	"book/helper"
	"book/pb/book"
	"context"
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (b *bookRepo) GetBooks(ctx context.Context, req *book.BookFindAllRequest) (res *book.BookFindAllResponse, err error) {
	s := req.Search
	// status := req.Status

	searchMatch := helper.GetSerchRegex([]string{
		"book_id",
		"author",
	}, s)

	filter := bson.M{
		"$or": searchMatch,
	}

	findOption := options.Find()

	page := req.Page
	if page == 0 {
		page = 1
	}
	perPage := req.PerPage
	if perPage == 0 {
		perPage = 10
	}
	offset := (page - 1) * perPage

	findOption.SetSkip(offset)
	findOption.SetLimit(perPage)

	cur, err := b.book.Find(ctx, filter)
	if err != nil {
		return
	}

	var books = []*book.Book{}

	for cur.Next(ctx) {
		each := domain.Book{}
		err := cur.Decode(&each)
		if err != nil {
			return nil, err
		}

		item := b.parseBook(each)
		books = append(books, item)
	}

	dataSize := int64(len(books))
	if dataSize < 1 {
		res = &book.BookFindAllResponse{
			IsEmpty: true,
		}

		return
	}

	res = &book.BookFindAllResponse{Payload: &book.BookFindAll{
		Books: books,
	}}

	rows, _ := b.book.CountDocuments(ctx, filter)
	res.Payload.Rows = rows
	var pages int64 = 1
	if rows >= perPage {
		pages = int64(math.Ceil(float64(rows) / float64(perPage)))
	}
	res.Payload.Pages = pages
	res.Payload.PerPage = perPage
	res.Payload.ActivePage = page
	res.Payload.Total = dataSize

	return
}
