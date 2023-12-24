# Build stage
FROM golang:1.20-alpine as builder

WORKDIR /app

# Copy the entire project
COPY . .

# Set the Go environment variables
ENV GOOS=linux
ENV GOARCH=amd64

# Build the Go binary
RUN go build -o coordinator-exec cmd/coordinator/main.go

# Runtime stage
FROM alpine:3.19

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/coordinator-exec .

# Copy config and static/template directories
COPY --from=builder /app/cmd/coordinator/config.json .
COPY --from=builder /app/coordinator/static ./coordinator/static
COPY --from=builder /app/coordinator/templates ./coordinator/templates

EXPOSE 8080

CMD ["./coordinator-exec", "config.json"]
