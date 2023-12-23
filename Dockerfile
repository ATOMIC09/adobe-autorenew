# Start from a minimal Go image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY main.go .

# Copy the .env file into the container
COPY .env .

# Build the Go application
RUN go build -o myapp main.go

# Run the compiled application with secrets passed as arguments
CMD ["./myapp"]
