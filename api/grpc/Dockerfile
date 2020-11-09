# reference: 
# https://www.cloudreach.com/en/resources/blog/cts-build-golang-dockerfiles/

#BUILD
FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build/api/grpc
RUN go build main.go

#DEPLOY
FROM alpine

# Add docker-compose-wait tool 
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /app/wait
RUN chmod +x /app/wait

# Deploy
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/api/grpc/main /app/
WORKDIR /app
CMD [ "./main" ]
