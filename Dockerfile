FROM golang:1.8

MAINTAINER David Butler

WORKDIR $GOPATH/src/dbuts/KoffeeCupLogger 
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["./go/src/main"]
