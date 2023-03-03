package repository

import (
	"book/domain"
	"book/helper"
	"book/pb/book"
	"book/variable"
	"context"
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (b *bookRepo) GetBooks(ctx context.Context, req *book.BookFindAllRequest) (res *book.BookFindAllResponse, err error) {
	s := req.Search
	sort := -1
	if req.Sort == "asc" {
		sort = 1
	}
	status := req.Status
	if status == "" {
		status = "none"
	}

	// deleted status
	deleted := bson.M{"deleted_at": bson.M{"$eq": nil}}
	if status == variable.StatusDeleted {
		deleted = bson.M{"deleted_at": bson.M{"$ne": nil}}
	}

	// out of stock status
	oos := bson.M{}
	if status == variable.StatusOutOfStock {
		oos = bson.M{"qty": bson.M{"$lte": 0}}
	}

	willEmpty := bson.M{}
	dangerStock := 5
	if status == variable.WillEmpty {
		willEmpty = bson.M{"$and": []bson.M{
			{"qty": bson.M{"$lte": dangerStock}},
			{"qty": bson.M{"$gt": 0}},
		}}
	}

	// additonal filter
	filterData := []bson.M{
		deleted,
		oos,
		willEmpty,
	}

	searchMatch := helper.GetSerchRegex([]string{
		"book_id",
		"title",
		"author",
		"publisher",
		"location_id",
		"uploader.user_id",
		"uploader.user",
	}, s)

	filter := bson.M{
		"$or":  searchMatch,
		"$and": filterData,
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
	findOption.SetSort(bson.M{"book_id": sort})

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
