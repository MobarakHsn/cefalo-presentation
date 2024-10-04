# Start from the official Golang base image
FROM golang:alpine AS builder


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o server .

# Start a new stage from a smaller base image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Install curl
RUN apk add --no-cache curl

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/server .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./server"]
