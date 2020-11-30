FROM golang:1.14-alpine3.12

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build api/main.go

# Export necessary port
EXPOSE 8080

CMD ["./main"]