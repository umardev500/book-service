run:
	go run main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	$(shell find ./pb -name "*.proto")

clean:
	find pb -name "*.pb.go" -type f -delete


book:
	grpcurl --plaintext -d '{"book_id": "1667832878734"}' localhost:5014 BookService.GetBook

books:
	grpcurl --plaintext -d '{"search": "721872389233", "status": "willEmpty"}' localhost:5014 BookService.GetBooks

delete:
	grpcurl --plaintext -d '{"book_id": "1667832878734"}' localhost:5014 BookService.Delete

softdelete:
	grpcurl --plaintext -d '{"book_id": "1667832878734"}' localhost:5014 BookService.SoftDelete

update:
	grpcurl --plaintext -d '{"book_id": "1667832878734", "payload": {"title": "Lord of The Ring", "uploader": {"user": "steven"}, "editors": [{"user_id": "721872389238", "user": "mark"}]}}' localhost:5014 BookService.Update

create:
	grpcurl --plaintext -d '{"book_id": "1667832878734","title": "Lord of The Ring","author": "J. R. R. Tolkien","publisher": "Allen & Unwin","pages_total": 450,"qty": 40,"cover": "cover.png","description": "The book description","location_id": "16673288374","uploader": {"user_id": "721872389233","user": "seteven"}}' localhost:5014 BookService.Create