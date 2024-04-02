# Stage 1: Build the application
FROM golang:1.22.1-alpine3.19 AS builder

WORKDIR /apiServer
COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go install github.com/codegangsta/gin@latest
COPY . .
# DEVELOPMENT SERVER
CMD ["gin", "run", "-i", "main.go", "--", "-p", "8000", "-a", "0.0.0.0"]

# FOR PRODUCTION BUILD ONLY
# Uncomment Below Lines for multi stage build
# RUN go build -o main

# # Stage 2: Create a minimal image with only the necessary artifacts
# FROM alpine:latest

# WORKDIR /app

# COPY --from=builder /apiServer/main .
# COPY --from=builder /apiServer/.env .


# CMD ["./main"]