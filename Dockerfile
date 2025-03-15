FROM golang:1.24 as base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

RUN go install github.com/air-verse/air@latest

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]