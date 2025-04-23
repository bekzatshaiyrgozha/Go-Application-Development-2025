FROM golang:1.22-alpine

# Ограничения не задаются непосредственно в Dockerfile
WORKDIR /app
COPY . .
RUN go build -o main .

CMD ["./main"]