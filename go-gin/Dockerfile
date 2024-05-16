# Use the official Golang image to create a build environment.
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container.
WORKDIR /app

# Copy go mod and sum files.
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
RUN go mod download

# Copy the source code into the container's workspace.
COPY . .

# Build the command inside the container.
# Compile the binary statically, including all dependencies.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Use a Docker multi-stage build to create a lean production image.
FROM alpine:latest  

WORKDIR /root/

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/myapp .

EXPOSE 5000

CMD ["./myapp"]
