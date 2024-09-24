FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY internal internal

COPY cmd cmd

RUN CGO_ENABLED=0 GOOS=linux go build -o key_generation_service ./cmd/key_generation_service

FROM alpine:3.19 AS security_provider

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

FROM scratch AS runtime

COPY --from=security_provider /etc/passwd /etc/passwd

USER nonroot

COPY --from=builder /app/key_generation_service /app/key_generation_service

EXPOSE 8081

ENTRYPOINT ["/app/key_generation_service"]

