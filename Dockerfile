# Start with the official Golang image as the build stage
FROM golang:1.24 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files first and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# Use a minimal base image for the final container
FROM alpine:latest  

# Set the working directory inside the container
WORKDIR /usr/local/bin

# Copy the compiled binary from the build stage
COPY --from=build /app/app .

# Expose the application's port (change according to your app)
EXPOSE 8080

# Command to run the application
CMD ["app"]

