FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -ldflags="-s -w" -o app ./cmd/gotime/main.go

FROM alpine:latest

RUN apk --no-cache add tzdata

WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 3000

USER 3000:3000

CMD ["./app"]
