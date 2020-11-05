# reference: 
# https://www.cloudreach.com/en/resources/blog/cts-build-golang-dockerfiles/

#BUILD
FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build/server/json_rpc_http
RUN go build server.go

#DEPLOY
FROM alpine

# Add docker-compose-wait tool 
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /app/wait
RUN chmod +x /app/wait

# Deploy
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/server/json_rpc_http/server /app/
WORKDIR /app
CMD [ "./server" ]
