# Use a Golang base image to build the application
FROM golang:1.22.2 AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main

# Use a minimal base image with required GLIBC version
FROM debian:bookworm-slim

# Install necessary dependencies
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Copy the config file to the working directory
COPY config.yaml /app/config.yaml

# Expose the port the application will run on
EXPOSE 8080

# Run the application
CMD ["./main"]
