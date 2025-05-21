FROM golang:1.24.2 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./app/cmd/webserver

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
RUN apk add --no-cache file && ls -l /app && file /app/app
CMD ["./app"]