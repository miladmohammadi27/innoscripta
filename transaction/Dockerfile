# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o transaction main.go

# Stage 2: Create a smaller image for the final application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/transaction .

# Run the Go application
CMD ["./transaction", "serve"]
