FROM golang:1.9.2-alpine

WORKDIR /go/src/github.com/antonybudianto/go-starter

RUN apk --update add git
RUN go-wrapper download -u github.com/golang/dep/cmd/dep \
    && go-wrapper install github.com/golang/dep/cmd/dep

COPY . .

RUN dep ensure

CMD go run main.go
