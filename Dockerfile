FROM golang:1.14.15-alpine3.11 as builder

RUN mkdir -p /app
WORKDIR /app

COPY go.mod go.sum ./

RUN set -x; go mod download

COPY . .

RUN set -x; go build -o linux-audit-exporter

FROM alpine:3.11

COPY --from=builder /app/linux-audit-exporter /usr/local/bin/

CMD ["/usr/local/bin/linux-audit-exporter"]
