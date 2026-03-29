FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8082
CMD ["./main"]
