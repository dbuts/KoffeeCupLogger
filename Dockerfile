FROM golang:1.8

MAINTAINER David Butler

RUN mkdir /app
ADD . /app

WORKDIR /app

RUN go get -d -v ./...

EXPOSE 8080

RUN go build -o main ./go/src

CMD ["/app/main"]
