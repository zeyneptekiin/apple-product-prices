FROM golang:1.22

WORKDIR /app

COPY src/go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main"]
