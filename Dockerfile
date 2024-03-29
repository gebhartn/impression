## Build
FROM golang:1.17-alpine3.14 AS builder

RUN apk add --update gcc musl-dev ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir -p /app
ADD . /app
WORKDIR /app

RUN adduser -u 10001 -D app

RUN GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=1 \
    go build -ldflags='-extldflags=-static' -o app .

RUN chown app: ./database

## Copy
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
USER app

WORKDIR /app

COPY --from=builder /app/app ./app
COPY --from=builder /app/database ./database
VOLUME ./database

## Run
CMD ["./app"]
