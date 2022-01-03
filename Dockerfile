FROM golang:1.17 AS builder

WORKDIR /app

COPY . .

RUN go get -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server/main.go

FROM alpine:latest AS runner

WORKDIR /app

COPY --from=builder /app/server .

COPY --from=builder /app/.env .

CMD ["./server"]