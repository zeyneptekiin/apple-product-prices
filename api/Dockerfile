FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./src/ ./src/

RUN go build -o /app/main ./src

CMD ["/app/main"]
