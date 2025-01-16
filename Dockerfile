FROM ubuntu:latest
LABEL authors="BUITOFU"


FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY .env .
COPY go.mod .

RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build -o /app/main .

FROM alpine:latest

WORKDIR /root/

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 6969

CMD ["./main"]


