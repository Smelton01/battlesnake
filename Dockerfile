##
## Build base binary
##
FROM golang:1.18-alpine AS base

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download


# Copy the source from the current directory to the Working Directory inside the container
COPY . /app/


# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main ./cmd/server/main.go

##
## Deploy 
##
### make this alpine
FROM alpine:3.14

WORKDIR /

COPY --from=base /main /main

# RUN apk --no-cache add ca-certificates
# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Expose port 80 to the outside world
EXPOSE 80

# Run the executable
ENTRYPOINT ["/main"]