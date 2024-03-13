FROM golang:1.22.1-alpine AS builder

COPY . /github.com/croked91/chat-server
WORKDIR /github.com/croked91/chat-server

RUN go mod download
RUN go build -o ./bin/chat-server internal/cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/croked91/chat-server/bin/chat-server .

CMD ["./chat-server"]