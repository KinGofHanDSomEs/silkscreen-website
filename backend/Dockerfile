FROM golang:alpine AS builder
WORKDIR /app

COPY ../../learning/go/crud/backend/go.mod go.sum ./
RUN go mod download

COPY ../../learning/go/crud/backend .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]