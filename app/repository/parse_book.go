package repository

import (
	"book/domain"
	"book/pb/book"
)

func (b *bookRepo) parseBook(item domain.Book) (res *book.Book) {

	uploader := &book.BookUploader{
		UserId: item.Uploader.UserId,
		User:   item.Uploader.User,
	}

	var editors []*book.BookEditor
	if item.Editors != nil && len(item.Editors) > 0 {
		for _, val := range item.Editors {
			foo := book.BookEditor{UserId: val.UserId, User: val.User}
			editors = append(editors, &foo)
		}
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
		Editors:     editors,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		DeletedAt:   item.DeletedAt,
	}

	return
}
