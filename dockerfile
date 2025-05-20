FROM golang:1.24.2 as builder

WORKDIR /app
COPY . .
RUN go build -o app ./cmd/webserver
WORKDIR /app/cmd/webserver

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]