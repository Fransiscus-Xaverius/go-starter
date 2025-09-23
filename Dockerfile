# Stage 1: Build the Go binary
FROM golang:1.25-alpine AS builder


# Install Git
RUN apk update && apk add --no-cache git coreutils openssh-client

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /my-app .

# Stage 2: Final minimal runtime image
FROM alpine:latest

# Install only the minimal runtime dependencies (if any required, e.g., libc)
RUN apk --no-cache add ca-certificates tzdata &&  \
    ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

# Copy the compiled Go binary from the builder stage
COPY --from=builder /my-app /my-app

# Set the command to run the Go application
CMD ["/my-app"]