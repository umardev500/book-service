package repository

import (
	"book/domain"
	"book/pb/book"
)

func (b *bookRepo) parseBook(item domain.Book) (res *book.Book) {
	uploader := &book.BookUploader{
		UserId: item.Uploader.UserId,
	}

	res = &book.Book{
		BookId:      item.BookId,
		Title:       item.Title,
		Author:      item.Author,
		Publisher:   item.Publisher,
		PagesTotal:  item.PagesTotal,
		Qty:         item.Qty,
		Cover:       item.Cover,
		Description: item.Description,
		LocationId:  item.LocationId,
		Uploader:    uploader,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		DeletedAt:   item.DeletedAt,
	}

	return
}
