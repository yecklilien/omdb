package main

import (
	"fmt"
	logger "log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/yecklilien/OMDB/api/handler"
	"github.com/yecklilien/OMDB/repository"
	"github.com/yecklilien/OMDB/usecase/log"
	"github.com/yecklilien/OMDB/usecase/movie"
	
)

func main() {

	db, err := gorm.Open(mysql.Open(constructDSN()), &gorm.Config{})
	if err != nil {
		logger.Fatalf("failed to connect DB: %v with DSN:%v", err, constructDSN())
	}

	movieService := movie.NewService(os.Getenv("OMDB_API_KEY"))

	logRepo := repository.NewLog(db)

	logService := log.NewService(logRepo)

	handler := handler.NewJSONHTTPHandler(movieService, logService)

	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	server.RegisterService(handler, "rpc")

	router := mux.NewRouter()
	router.Handle("/", server)
	err = http.ListenAndServe(":9002", router)
}

func constructDSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))
}
