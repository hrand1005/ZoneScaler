# Build stage
FROM golang:1.20 as builder

WORKDIR /app

COPY . .

# Set the Go environment variables
ENV GOOS=linux
ENV GOARCH=amd64

# Copy the entire project
# Ensure that the Docker context is set to the root of your Go project

# Build the Go binary and gather necessary files
RUN mkdir -p /app/tmp/coordinator/static && \
    mkdir -p /app/tmp/coordinator/template && \
    go build -o /app/tmp/coordinator-exec cmd/coordinator/main.go && \
    cp cmd/coordinator/config.json /app/tmp/config.json && \
    cp -r coordinator/static/ /app/tmp/coordinator/static && \
    cp -r coordinator/templates/ /app/tmp/coordinator/templates
# Runtime stage
FROM alpine:3.19

WORKDIR /app

# Copy only the necessary artifacts from the builder stage
COPY --from=builder /app/tmp/coordinator-exec .
COPY --from=builder /app/tmp/config.json .

# Copy static and template directories into a coordinator directory
RUN mkdir coordinator
COPY --from=builder /app/tmp/coordinator/static ./coordinator/static
COPY --from=builder /app/tmp/coordinator/templates ./coordinator/templates

EXPOSE 8080

CMD ["./coordinator-exec", "config.json"]
