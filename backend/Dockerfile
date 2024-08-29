# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o pennywise-backend ./cmd/main.go

# Stage 2: Run the Go binary in a smaller image
FROM alpine:latest

WORKDIR /app

# Copy the Pre-built binary file from the builder stage
COPY --from=builder /app/pennywise-backend .

# Install Go in the final image to run the script
RUN apk add --no-cache go

# Copy the source code and scripts to the final image
COPY . .

# Create the .env.dev file using environment variables passed from Docker Compose
RUN touch .env.dev

# Expose the port the service runs on
EXPOSE 8080

# Command to run the script to generate JWT secret and then start the application
CMD ["/bin/sh", "-c", "echo 'DB_USER=${DB_USER}' >> .env.dev && \
    echo 'DB_PASSWORD=${DB_PASSWORD}' >> .env.dev && \
    echo 'DB_HOST=${DB_HOST}' >> .env.dev && \
    echo 'DB_PORT=${DB_PORT}' >> .env.dev && \
    echo 'DB_NAME=${DB_NAME}' >> .env.dev && \
    go run ./scripts/generate_secret.go -env .env.dev && \
    ./pennywise-backend"]
