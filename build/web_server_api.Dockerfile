FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY internal internal

COPY cmd cmd

RUN CGO_ENABLED=0 GOOS=linux go build -o web_server_api ./cmd/web_server_api

FROM alpine:3.19 AS security_provider

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

FROM scratch AS runtime

COPY --from=security_provider /etc/passwd /etc/passwd

USER nonroot

COPY --from=builder /app/web_server_api /app/web_server_api

EXPOSE 8080

ENTRYPOINT ["/app/web_server_api"]
