# omdb

## How to run the server
Make sure you already had [docker] installed.

```
Fill .env file with OMDB API KEY and MY SQL configuration
Then run below command
sudo docker compose up
```

Docker will spawn container for each grpc and json http server
And spawn mysql server (to store the log)

## Test json http over rpc

```
curl -X POST \
  http://localhost:9002 \
  -H 'Accept: */*' \
  -H 'Content-Type: application/json; charset=utf-8'
  -d '{
    "jsonrpc": "2.0",
    "method": "rpc.SearchMovie",
    "id": "1",
    "params": [
    	{
    		"Page" : 1,
    		"Query" : "Batman"
    	}
    ]
}'
```

```
curl -X POST \
  http://localhost:9002 \
  -H 'Accept: */*' \
  -H 'Content-Type: application/json; charset=utf-8'
  -d '{
    "jsonrpc": "2.0",
    "method": "rpc.GetMovieDetail",
    "id": "1",
    "params": [
    	{
    		"imdbID": "tt0052602"
    	}
    ]
}'
```

## Test GRPC
```
cd cmd/client/
go run client_grpc.go
```

## Reference

List of reference im made to create this

https://www.cloudreach.com/en/resources/blog/cts-build-golang-dockerfiles/

https://docs.docker.com/compose/gettingstarted/

https://medium.com/better-programming/using-variables-in-docker-compose-265a604c2006

https://gorm.io/index.html

https://godoc.org/github.com/gorilla/rpc


[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

[docker]: https://docs.docker.com/get-docker/