FROM golang:1.24.2 as builder

WORKDIR /app
COPY . .
RUN go build -o app ./cmd/werbserver
WORKDIR /app/cmd/werbserver

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]