FROM golang:1.22.1-alpine3.19

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8080

RUN go build -o nytaxi cmd/api/main.go

CMD ["/app/nytaxi"]

