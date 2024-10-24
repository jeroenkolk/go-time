FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app ./cmd/gotime/main.go

FROM alpine:latest

RUN apk --no-cache add tzdata

WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 3000

USER nonroot:nonroot

CMD ["./app"]