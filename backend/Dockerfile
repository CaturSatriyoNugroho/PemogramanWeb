# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install dependencies first
COPY go.mod ./
RUN go mod tidy

# Copy the rest of the code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/api main.go

# Run stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api"]
