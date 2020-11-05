package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/yecklilien/OMDB/impl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(mysql.Open(constructDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect DB: %v with DSN:%v", err, constructDSN())
		os.Exit(1)
	}

	movieService := impl.NewMovieAPI(db, os.Getenv("OMDB_API_KEY"))

	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	rpcService := impl.NewMovieAPIHTTPJSONRPCServer(movieService)

	server.RegisterService(rpcService, "rpc")

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
