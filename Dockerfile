# Stage ##############################################################
# FROM base AS builder
FROM golang:alpine AS build

# Set the working directory
WORKDIR /insights

# Copy and download dependencies
COPY go.mod  ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o appbin booking-req-insights/cmd/insights

# ##############################################################################
# # Release Stage ##############################################################
FROM alpine:edge

# Set the working directory
WORKDIR /insights

# Copy the binary from the build stage and the .env file
COPY --from=build /insights/appbin .

# TODO: pass the envs
# COPY --from=build /insights/.env

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata

# Set the entrypoint command
ENTRYPOINT ["/app/myapp"]


