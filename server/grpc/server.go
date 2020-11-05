package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/yecklilien/OMDB/impl"
	"github.com/yecklilien/OMDB/movie"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(mysql.Open(constructDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect DB: %v with DSN:%v", err, constructDSN())
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}

	movieService := impl.NewMovieAPI(db, os.Getenv("OMDB_API_KEY"))

	server := grpc.NewServer()

	grpcService := impl.NewMovieAPIGRPCServer(movieService)

	movie.RegisterMovieAPIServer(server, grpcService)

	err = server.Serve(lis)
}

func constructDSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))
}
