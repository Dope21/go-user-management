FROM golang:1.24 as base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM base AS dev
RUN go install github.com/air-verse/air@latest
EXPOSE 8080
CMD ["air", "-c", ".air.toml"]

FROM golang:1.24-alpine AS prod
WORKDIR /app
COPY --from=base /app/main .
COPY --from=base /app/.env .
COPY --from=base /app/db ./db
EXPOSE 8080
CMD ["./main"]