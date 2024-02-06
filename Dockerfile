FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="rc <richard.cargill@pm.me>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
# COPY go.mod go.sum ./
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and the
# go.sum files are not changed 
RUN go mod download 

COPY . .

# TODO
# Check what all this stuff does
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shorty .

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./shorty"]
