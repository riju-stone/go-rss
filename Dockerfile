FROM golang:1.22

WORKDIR /go-rss-backend

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . ./

RUN go build -o bin/go-rss main.go

EXPOSE 8000

CMD ["bin/go-rss"]
