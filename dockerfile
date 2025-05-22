FROM golang:1.24.2 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o webserver ./app/cmd/webserver

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/webserver .
COPY /cmd/webserver/game.db.json .
EXPOSE 8080
CMD ["./webserver"]