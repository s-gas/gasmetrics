# 1
FROM golang:1.26 AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o gasmetrics-bot ./cmd/

# 2
FROM debian:12.13

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/gasmetrics .

CMD ["./gasmetrics-bot"]
