FROM golang:1.18 as builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy the Go Modules manifests and download any dependencies.
# These layers are only re-built when go.mod or go.sum files change.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code.
COPY . .

# Build the Go app.
# Compile the binary statically for compatibility with the scratch image.
# Make sure to specify the path to the directory containing your main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Use a minimal runtime image.
FROM alpine:latest  

# Add CA certificates and other dependencies including Dockerize
RUN apk update && apk add --no-cache ca-certificates bash wget \
    && update-ca-certificates

# Install Dockerize
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /app/main .

# Expose port 8080 (if your app uses a different port, update it here).
EXPOSE 8080

# Command to run the executable.
CMD ["./main"]
