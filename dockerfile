# Stage 1: Build Stage
FROM golang:1.22 AS builder

WORKDIR /app

RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Runtime Stage
FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/main .
COPY data.json ./data.json

EXPOSE 8080

CMD ["./main"]
