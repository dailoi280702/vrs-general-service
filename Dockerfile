FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o vrs-general-service main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/vrs-general-service .

EXPOSE 9001

CMD ["./vrs-general-service"]
