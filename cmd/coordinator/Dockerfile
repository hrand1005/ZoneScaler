# Build stage
FROM golang:1.20 as builder

WORKDIR /app

COPY . .

# Set the Go environment variables
ENV GOOS=linux
ENV GOARCH=amd64

# Build the Go binary and gather necessary files
RUN mkdir -p /app/tmp && \
    go build -o /app/tmp/coordinator cmd/coordinator/main.go && \
    cp cmd/coordinator/config.json /app/tmp/config.json

# Runtime stage
FROM alpine:3.19

WORKDIR /app

# Copy only the necessary artifacts from the builder stage
COPY --from=builder /app/tmp/coordinator .
COPY --from=builder /app/tmp/config.json .

EXPOSE 8080

CMD ["./coordinator", "config.json"]
