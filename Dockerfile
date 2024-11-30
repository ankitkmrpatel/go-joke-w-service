# Use the latest golang image as the base image
FROM golang:1.23.3-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main.go

# Expose the metrics port (Prometheus will scrape this)
EXPOSE 9090

# Command to run the application
CMD ["./main"]
