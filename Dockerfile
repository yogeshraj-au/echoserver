# Use a small base image like Alpine for the final image
FROM golang:1.18 AS builder

# Set the working directory in the container
WORKDIR /go/src/app

# Copy the source code into the container
COPY . .

# Build the echoserver binary
RUN go build -o /echoserver .

# Start a new stage from the same base image
FROM golang:1.18

# Copy the built echoserver binary from the builder stage
COPY --from=builder /echoserver /echoserver

# Expose the port for the echoserver
EXPOSE 8080

# Command to run the echoserver
CMD ["/echoserver"]
