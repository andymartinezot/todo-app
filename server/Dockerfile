# Use the official Golang image as a base image
FROM golang:1.18

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# List files to debug
RUN ls -la /app

# Expose port 9000 to the outside world
EXPOSE 9000

# Ensure the binary is executable
RUN chmod +x ./main

# Command to run the executable
CMD ["go", "run", "main.go"]
#CMD ["./main"]