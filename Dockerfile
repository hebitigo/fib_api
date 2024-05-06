FROM golang:1.22.1-alpine3.18

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o api .

EXPOSE 8080

CMD ["./api"]