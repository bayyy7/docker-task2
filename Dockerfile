# Start from golang base image
FROM golang:1.23.2-alpine AS builder

# Set working directory
WORKDIR /app

# Copy all files
COPY . .

# Build the application
RUN go build -o main .

# Start a new stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy AUTHORS.md
COPY AUTHORS.md .

# Expose port 80
EXPOSE 80

# Command to run the executable
CMD ["./main"]