FROM golang:1.20.3-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o server cmd/server.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/server .
COPY .env .
COPY database/migrations ./database/migrations

CMD [ "/app/server" ]
