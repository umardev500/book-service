package domain

import (
	"book/pb"
	"book/pb/book"
	"context"
)

type BookUploader struct {
	UserId string `bson:"user_id"`
	User   string `bson:"user"`
}

type BookEditor struct {
	UserId string `bson:"user_id"`
	User   string `bson:"user"`
}

type Book struct {
	BookId      string       `bson:"book_id"`
	Title       string       `bson:"title"`
	Author      string       `bson:"author"`
	Publisher   string       `bson:"publisher"`
	PagesTotal  int64        `bson:"pages_total"`
	Qty         int64        `bson:"qty"`
	Cover       string       `bson:"cover"`
	Description string       `bson:"description"`
	LocationId  string       `bson:"location_id"`
	Uploader    BookUploader `bson:"uploader"`
	Editors     []BookEditor `bson:"editors"`
	CreatedAt   int64        `bson:"created_at"`
	UpdatedAt   int64        `bson:"updated_at"`
	DeletedAt   int64        `bson:"deleted_at"`
}

type BookService interface {
	Create(ctx context.Context, req *book.BookCreateRequest) (res *book.BookCreateResponse, err error)
	GetBook(ctx context.Context, req *book.BookFindOneRequest) (res *book.BookFindOneResponse, err error)
	GetBooks(ctx context.Context, req *book.BookFindAllRequest) (res *book.BookFindAllResponse, err error)
	Delete(ctx context.Context, req *book.BookDeleteRequest) (res *pb.OperationResponse, err error)
	SoftDelete(ctx context.Context, req *book.BookDeleteRequest) (res *pb.OperationResponse, err error)
	Update(ctx context.Context, req *book.BookUpdateRequest) (res *pb.OperationResponse, err error)
	Restore(ctx context.Context, req *book.BookRestoreRequest) (res *pb.OperationResponse, err error)
}

type BookRepo interface {
	Create(ctx context.Context, req *book.BookCreateRequest) (res *book.BookCreateResponse, err error)
	GetBook(ctx context.Context, req *book.BookFindOneRequest) (res *book.BookFindOneResponse, err error)
	GetBooks(ctx context.Context, req *book.BookFindAllRequest) (res *book.BookFindAllResponse, err error)
	Delete(ctx context.Context, req *book.BookDeleteRequest) (res *pb.OperationResponse, err error)
	SoftDelete(ctx context.Context, req *book.BookDeleteRequest) (res *pb.OperationResponse, err error)
	Update(ctx context.Context, req *book.BookUpdateRequest) (res *pb.OperationResponse, err error)
	Restore(ctx context.Context, req *book.BookRestoreRequest) (res *pb.OperationResponse, err error)
}
