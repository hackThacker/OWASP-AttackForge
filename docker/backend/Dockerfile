# ==============================================================================
# Build Stage
# ==============================================================================
FROM golang:alpine AS builder

# Install required certificates and build tools
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy dependency files first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Compile the Go application statically for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -trimpath -ldflags="-w -s" -o server ./cmd/server

# ==============================================================================
# Final Stage
# ==============================================================================
FROM scratch

# Copy static assets and built binary
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/server /server

# Expose ports: 8081 (REST/WS), 50051 (gRPC)
EXPOSE 8081 50051

# Run the backend service
ENTRYPOINT ["/server"]
