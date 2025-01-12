# Build stage
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the core service
RUN go build -o core ./cmd/core/main.go

# Build the notification service
RUN go build -o notification ./cmd/notification/main.go

# Final stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executables from the builder stage
COPY --from=builder /app/core .
COPY --from=builder /app/notification .

# Expose the ports the apps run on
EXPOSE 8080
EXPOSE 9090

# Command to run both executables
CMD ["sh", "-c", "./core & ./notification"]
