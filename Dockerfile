# Stage ##############################################################
# FROM base AS builder
FROM golang:alpine AS build

# Set the working directory
WORKDIR /booking-insights

# Copy and download dependencies
COPY go.mod  ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o http-server-api /booking-insights/cmd/http

# ##############################################################################
# # Release Stage ##############################################################
FROM alpine:edge

# Set the working directory
WORKDIR /booking-insights

# Copy the binary from the build stage
COPY --from=build /booking-insights .

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata

# Set the entrypoint command
ENTRYPOINT ["/booking-insights/http-server-api"]


