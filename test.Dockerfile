# Start from a base Go image
FROM golang:1.22
# Set the working directory inside the container
WORKDIR /app

# Copy the entire project into the container
COPY . .

# Install any Go modules necessary for the project
RUN go mod download

# Build the main Go application
RUN go build -o test /app/f1/main.go

# Default command to run when the container starts (can be overridden)
ENTRYPOINT ["/app/test"]
