# Use the official Golang image as a builder stage
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal Alpine Linux image as the final stage
FROM alpine:latest

# Set the working directory inside the final container
WORKDIR /app

# Copy the compiled binary from the builder stage to the final container
COPY --from=builder /app/main .

# Expose port 8090 for the server to listen on
EXPOSE 8090

# Run the Go application
CMD ["./main", "serve", "--http=0.0.0.0:8090"]
