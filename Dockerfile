# Use the official Go image as a builder stage
FROM golang:1.21 as builder

RUN mkdir /app

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . ./

# Build the Go application
RUN go build -o /young-astrologer-service/

# Use a smaller base image for the final stage
FROM alpine:latest

# Copy the binary from the builder stage into the final image
COPY --from=builder /young-astrologer-service .

# Expose the port your application listens on (if needed)
EXPOSE 8080

# Define the command to run your application
CMD ["./young-astrologer-service"]
