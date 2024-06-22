# Use the official Golang image to create a build artifact.
# This is the base image for building the Go application.
FROM golang:1.20-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/server

# Use a smaller image for running the app
FROM alpine:latest

# Install necessary packages
RUN apk add --no-cache ca-certificates

# Copy the compiled Go binary into the smaller image
COPY --from=builder /app/main /app/main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/app/main"]
