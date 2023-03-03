package main

import (
	"book/config"
	"book/injector"
	"book/pb/book"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	db := config.NewConn().Database("db_book")
	bookCollection := db.Collection("books")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	bookInjector := injector.NewBookInjector(bookCollection)

	book.RegisterBookServiceServer(server, bookInjector)

	fmt.Printf("⚡️[server]: gRPC Server is running on port %s\n", port)
	log.Fatal(server.Serve(lis))
}
