# Specifies a parent image
FROM golang:1.23.1-alpine AS builder

# Creates an app directory to hold your app’s source code
WORKDIR /app

# Copies everything from your root directory into /app
COPY . .

# Installs Go dependencies
RUN go mod download

# Build the Go app
RUN go build -o /godocker

FROM scratch
COPY --from=builder /godocker /
# Tells Docker which network port your container listens on
# Builds your app with optional configuration
# Set the environment variable for the build
ENV GIN_MODE=release
EXPOSE 8080
CMD ["/godocker"]
