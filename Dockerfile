# Use the official golang image to create a build artifact.
FROM golang:1.21.1 as builder

# Create and set the working directory.
WORKDIR /filmgrpc

# Copy the Go Modules manifests.
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
RUN go mod download

# Copy the local package files to the container's workspace.
COPY . .

# Build the binary for the server.
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./service

# Build the binary for the client.
RUN CGO_ENABLED=0 GOOS=linux go build -o client ./client

# Use alpine image for a minimal docker image
FROM alpine:latest

# Set the working directory in the container.
WORKDIR /app

# Copy the pre-built binary from the builder stage to the /app directory in the container.
COPY --from=builder /filmgrpc/server /app/server
COPY --from=builder /filmgrpc/client /app/client


# Expose the port on which the server will run.
EXPOSE 50051

# Command to run the server when the container starts.
CMD ["/app/server"]