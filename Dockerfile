# Building stage
FROM golang:1.20.0-alpine3.17 AS builder

# maintainer info
LABEL maintainer = "Abhijith A <abhijithworkemail@gmail.com>"

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main

# Final running stage
FROM alpine:latest

WORKDIR /app

COPY .env .

COPY --from=builder /app/main .

RUN chmod +x ./main

CMD ["./main"]

EXPOSE 3000

