# Stage 1: Build the Go binary
FROM golang:1.24.6 AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# tells docker to load every variable from the file to the containers environment
COPY .env /root/.env

# Copy the rest of the source code
COPY . .

# Build the binary (statically linked)
RUN CGO_ENABLED=0 GOOS=linux go build -o smart-notes-api ./cmd/server

# Stage 2: Minimal final image
FROM scratch

# Copy the binary from builder
COPY --from=builder /app/smart-notes-api /smart-notes-api

# Expose app port
EXPOSE 8080

# Start the app
ENTRYPOINT ["/smart-notes-api"]
