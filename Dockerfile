FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

COPY . .

ARG TARGETARCH
RUN mkdir generated && oapi-codegen --config=config.yaml api/openapi.yaml
RUN CGO_ENABLED=1 GOOS=linux GOARCH=$TARGETARCH go build -ldflags="-s -w" -o app ./cmd/gotime/main.go

FROM alpine:latest

RUN apk --no-cache add tzdata

WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 3000

USER 3000:3000

CMD ["./app"]
