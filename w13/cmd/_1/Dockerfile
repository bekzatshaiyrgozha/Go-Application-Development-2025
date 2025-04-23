## Base image with Go installed, using Alpine for a smaller image
FROM golang:1.22-alpine AS builder 

## Set the working directory in the container
WORKDIR /app

## Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./
RUN go mod download


## Copy the rest of the application code to the working directory
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

## Create a smaller image for the final build with no Go dependencies
FROM alpine:latest

## Set the working directory in the container again
WORKDIR /root/
COPY --from=builder /app/main .

## Expose the port the app runs on
EXPOSE 8080

## Command to run the application
CMD ["./main"]