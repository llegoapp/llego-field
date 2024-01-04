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
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a minimal runtime image.
FROM alpine:latest  

# Add CA certificates for HTTPS connections.
RUN apk --no-cache add ca-certificates

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /app/main .

# Expose port 8080 (if your app uses a different port, update it here).
EXPOSE 8080

# Command to run the executable.
CMD ["./main"]
