# Base image
FROM golang:1.19.3-alpine AS builder

# Create workspace
WORKDIR /opt/app/

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Fetch dependencies
RUN go mod download && \
    go mod verify

# Copy the source code as the last step
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/microservicio-media-api ./cmd/app

# Create uploads directory
RUN mkdir /uploads

# Build a small image
FROM alpine:3.17.3

# Copy user and group files from the builder
COPY --from=builder /etc/passwd /etc/passwd

# Copy the static executable
COPY --from=builder /go/bin/microservicio-media-api /microservicio-media-api

# Run as root user
USER root

# Run app
ENTRYPOINT ["./microservicio-media-api"]
