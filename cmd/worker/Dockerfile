FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the entire worker application
COPY ../../ .

# Build the worker application
RUN go build -o wrk ./cmd/worker

# Expose any necessary ports
EXPOSE 8081

# Command to run the application
ENTRYPOINT [ "./wrk", "./cmd/worker/config.json" ]
