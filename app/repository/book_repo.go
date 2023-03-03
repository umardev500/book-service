package repository

import (
	"book/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type bookRepo struct {
	book *mongo.Collection
}

func NewBookRepo(book *mongo.Collection) domain.BookRepo {
	return &bookRepo{
		book: book,
	}
}
