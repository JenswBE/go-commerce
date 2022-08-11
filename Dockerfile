FROM golang:1.18 AS builder

WORKDIR /src/
COPY . .
RUN go install github.com/nishanths/exhaustive/...@latest
RUN exhaustive ./...
RUN CGO_ENABLED=0 go build -ldflags='-extldflags=-static' -o /bin/go-commerce-api

FROM alpine
COPY --from=builder /bin/go-commerce-api /go-commerce/bin/go-commerce-api
COPY --from=builder /src/docs/index.html /go-commerce/docs/index.html
COPY --from=builder /src/docs/openapi.yml /go-commerce/docs/openapi.yml
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
WORKDIR /go-commerce/bin
ENTRYPOINT ["./go-commerce-api"]
