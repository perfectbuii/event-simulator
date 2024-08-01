# Use the official Golang image as the base image
FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN GOPROXY="https://goproxy.io" go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main /app/f1/main.go

# Command to run the executable
CMD ["/app/main","run","constant","-r","1000/s","-d","2s","testAPIWithFranzKafka"]
