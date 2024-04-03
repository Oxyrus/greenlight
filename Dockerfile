# Use golang:alpine as the builder stage
FROM golang:alpine as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o greenlight ./cmd/api

# Use alpine:latest for the final stage
FROM alpine:latest

# Install necessary packages
RUN apk --no-cache add ca-certificates curl

# Set the working directory
WORKDIR /root

# Copy the migrations from the builder stage
COPY --from=builder /app/migrations ./migrations

# Download and install the migration tool
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz

# Copy the built application from the builder stage
COPY --from=builder /app/greenlight .

# Expose the necessary port
EXPOSE 4000

# Define the entrypoint script or command
CMD ./migrate -path ./migrations -database "postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:5432/${DB_NAME}?sslmode=disable" -verbose up \
    && ./greenlight -cors-trusted-origins="http://localhost:9000 http://localhost:9001"
