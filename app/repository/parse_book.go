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

	var editor []*book.BookEditor
	if item.Editor != nil && len(item.Editor) > 0 {
		for _, val := range item.Editor {
			foo := book.BookEditor{UserId: val.UserId, User: val.User}
			editor = append(editor, &foo)
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
		Editor:      editor,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		DeletedAt:   item.DeletedAt,
	}

	return
}
