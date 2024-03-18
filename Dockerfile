# Use the official Golang image to create a build artifact.
FROM golang:1.21 as builder

# Define arguments
ARG KEEPER_DB_PORT
ARG KEEPER_DB_HOST
ARG KEEPER_DB_USER
ARG KEEPER_DB_PASSWORD
ARG KEEPER_DB_NAME
ARG KEEPER_DB_SSLMODE
ARG KEEPER_SESSIONS_SECRET

# Set the environment variables from the arguments
ENV KEEPER_DB_PORT=${KEEPER_DB_PORT}
ENV KEEPER_DB_HOST=${KEEPER_DB_HOST}
ENV KEEPER_DB_USER=${KEEPER_DB_USER}
ENV KEEPER_DB_PASSWORD=${KEEPER_DB_PASSWORD}
ENV KEEPER_DB_NAME=${KEEPER_DB_NAME}
ENV KEEPER_DB_SSLMODE=${KEEPER_DB_SSLMODE}
ENV KEEPER_SESSIONS_SECRET=${KEEPER_SESSIONS_SECRET}

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o keeper .

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/keeper .

# Command to run the executable
CMD ["./keeper"]

# Expose connection port
EXPOSE 8888
