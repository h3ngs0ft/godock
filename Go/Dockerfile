FROM golang:latest

RUN apt-get update
RUN apt-get install nano -y
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/gorilla/mux"

EXPOSE 8080