package service

import "book/domain"

type bookService struct {
	repo domain.BookRepo
}

func NewBookService(repo domain.BookRepo) domain.BookService {
	return &bookService{repo: repo}
}
