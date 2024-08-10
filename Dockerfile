FROM golang:1.22.6

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/web/main.go

EXPOSE 8080

CMD ["./main"]