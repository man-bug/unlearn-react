# Start from the official Go image
FROM golang:1.22 as builder

# Set the working directory
WORKDIR /app

# Copy go mod file
COPY go.mod ./

# Copy go sum file if it exists
COPY go.sum* ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Install Tailwind CSS
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 \
    && chmod +x tailwindcss-linux-x64 \
    && mv tailwindcss-linux-x64 tailwindcss

# Build Tailwind CSS
RUN ./tailwindcss -i ./static/css/tailwind.css -o ./static/css/tailwind.output.css

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static
COPY --from=builder /app/internal/templates ./internal/templates

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable with logging
CMD ["/bin/sh", "-c", "echo 'Starting application...' && ./main 2>&1 | tee /var/log/app.log"]
