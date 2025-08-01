# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o gateway .

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/gateway .

EXPOSE 3000

CMD ["./gateway"]