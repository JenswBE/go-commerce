FROM golang:1.14-alpine AS builder

WORKDIR /src/
COPY . .
WORKDIR /src/api
RUN CGO_ENABLED=0 go build -ldflags='-extldflags=-static' -o /bin/go-commerce-api

FROM alpine
COPY --from=builder /bin/go-commerce-api /go-commerce/bin/go-commerce-api
COPY --from=builder /src/docs/index.html /go-commerce/docs/index.html
COPY --from=builder /src/docs/openapi.yml /go-commerce/docs/openapi.yml
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /go-commerce/bin
ENTRYPOINT ["./go-commerce-api"]