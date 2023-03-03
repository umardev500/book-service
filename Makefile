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