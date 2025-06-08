# Stage 1: build
FROM golang:1.24.4-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .

WORKDIR /app/services/auth-services
RUN go mod tidy
RUN go build -o auth-service ./cmd

# Stage 2: runtime
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/services/auth-services/auth-service .
COPY services/auth-services/.env .env

EXPOSE 8081
CMD ["./auth-service"]
