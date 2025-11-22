# Use golang docker image
FROM golang:1.25.4-alpine

# Sets a working directory
WORKDIR /app

# Install dependencies on alpine (gcc & musl-dev are required for sqlite3)
RUN apk add --no-cache git --update gcc musl-dev

# Copy go.mod and go.sum and download dependencies
COPY ../../go.mod ../../go.sum ./
RUN go mod download

# Copy the application source code
COPY ../.. .

# Build go app
RUN go build -o configs-service ./cmd/rest

# Expose app port
EXPOSE 8080

# Command for Run the Configs Service App inside the Container
CMD ["./configs-service"]