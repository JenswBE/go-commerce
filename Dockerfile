# Also update GitHub Actions workflow when bumping
FROM --platform=${BUILDPLATFORM} docker.io/library/golang:1.20 AS builder-base

FROM builder-base AS builder-amd64
ENV GOOS=linux
ENV GOARCH=amd64

FROM builder-base AS builder-armv6
ENV GOOS=linux
ENV GOARCH=arm
ENV GOARM=6

FROM builder-base AS builder-armv7
ENV GOOS=linux
ENV GOARCH=arm
ENV GOARM=7

FROM builder-base AS builder-arm64
ENV GOOS=linux
ENV GOARCH=arm64

FROM builder-${TARGETARCH}${TARGETVARIANT} AS builder
WORKDIR /src/
COPY . .
RUN GOARCH=amd64 go install golang.org/x/vuln/cmd/govulncheck@latest
RUN govulncheck ./...
RUN CGO_ENABLED=0 go build -ldflags='-extldflags=-static' -o /bin/go-commerce-api

FROM alpine
COPY --from=builder /bin/go-commerce-api /go-commerce/bin/go-commerce-api
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
WORKDIR /go-commerce/bin
ENTRYPOINT ["./go-commerce-api"]
