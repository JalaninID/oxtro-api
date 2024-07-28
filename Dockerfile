# Start from the latest golang base image
FROM golang:1.21.0-alpine AS builder

# Add Maintainer Info
LABEL maintainer="Triadmoko <triadmoko12@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go mod and sum files
COPY go.mod go.sum ./


# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# === Stage 2: Runtime Stage ===

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080
# environment variable
ENV DB_USER=postgres \
    DB_NAME=mydb \
    DB_PASS=postgres \
    DB_HOST=db \
    DB_PORT=5432 \
    DB_MIGRATION_PATH=./migration/ \
    JWT_TIME_DURATION=43800 \
    JWT_SECRET_KEY=NKy0VxAWfRW9ur0UE2DsSuc7eUJrU/Gdd7Qdk4YuImYEGCCvt2RkVAB8r4NWySuSjX2/ziPfE3A/Rwd5sJ0+uPMUsW/mlJA4Q8JTiUY783jxLdmZ5iG//qU/FQSuROqEbSjGVfPzfji+hc0A7S5Z+dE8hPo0DOW88VFdKwoUbzQ=

CMD ["./main"]
