package main

import (
	"fmt"
	logger "log"
	"net"
	"os"

	"github.com/yecklilien/OMDB/api/handler"
	"github.com/yecklilien/OMDB/api/presenter"
	"github.com/yecklilien/OMDB/repository"
	"github.com/yecklilien/OMDB/usecase/log"
	"github.com/yecklilien/OMDB/usecase/movie"

	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(mysql.Open(constructDSN()), &gorm.Config{})
	if err != nil {
		logger.Fatalf("failed to connect DB: %v with DSN:%v", err, constructDSN())
	}

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	movieService := movie.NewService(os.Getenv("OMDB_API_KEY"))

	logRepo := repository.NewLog(db)

	logService := log.NewService(logRepo)

	handler := handler.NewGRPCHandler(movieService, logService)

	server := grpc.NewServer()

	presenter.RegisterMovieAPIServer(server, handler)

	err = server.Serve(lis)
}

func constructDSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))
}
