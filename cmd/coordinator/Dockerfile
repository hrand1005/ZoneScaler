FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the entire coordinator application
COPY ../../ .

# Build the coordinator application
RUN go build -o crd ./cmd/coordinator

# Expose any necessary ports
EXPOSE 8080

# Command to run the application
ENTRYPOINT [ "./crd", "./cmd/coordinator/config.json" ]
