# Use a Go builder image
FROM golang:1.24.1 AS builder

WORKDIR /app

# Copy dependencies first (caching layer)
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the code and build the application
COPY . .
RUN go build -o main .  #  This ensures `main` is built

# Use a minimal image for running the app
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
