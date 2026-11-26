FROM golang:1.27.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o mypulse cmd/api/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/mypulse .

EXPOSE 8080

CMD ["./mypulse"]
