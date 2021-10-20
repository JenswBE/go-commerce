[![GitHub](https://img.shields.io/github/license/JenswBE/go-commerce)](https://github.com/JenswBE/go-commerce)
[![Docker Pulls](https://img.shields.io/docker/pulls/jenswbe/go-commerce)](https://hub.docker.com/r/jenswbe/go-commerce)
[![codecov](https://codecov.io/gh/JenswBE/go-commerce/branch/main/graph/badge.svg?token=S2oyV1sTWU)](https://codecov.io/gh/JenswBE/go-commerce)

# GoCommerce

KISS backend for an e-commerce

## Links

- GitHub: https://github.com/JenswBE/go-commerce
- DockerHub: https://hub.docker.com/r/jenswbe/go-commerce
- Admin UI: https://github.com/JenswBE/go-commerce-admin

## Configuration

GoCommerce can be configured in 2 ways:

1. Create a file called `config.yml` in the same folder or the parent folder of the binary. See `config.yml` for an example.
2. Set environment variables

If both are defined, the environment variables take precedence.

| Config key                | Env variable                | Description                                                                          | Default value  |
| ------------------------- | --------------------------- | ------------------------------------------------------------------------------------ | -------------- |
| Authentication.IssuerURL  | AUTH_ISSUER_URL             | URL to OpenID Configuration Issuer (without `.well-known/openid-configuration`)      |                |
| Database.Host             | DATABASE_HOST               | Hostname of the Postgres datatabase                                                  |                |
| Database.Port             | DATABASE_PORT               | Port of the Postgres datatabase                                                      | 5432           |
| Database.User             | DATABASE_USER               | Username for the Postgres datatabase                                                 |                |
| Database.Password         | DATABASE_PASSWORD           | Password for the Postgres datatabase                                                 |                |
| Database.Database         | DATABASE_DATABASE           | Database name for the Postgres datatabase                                            |                |
| ImageProxy.BaseURL        | IMAGE_PROXY_BASE_URL        | Base URL of your [Imgproxy instance](https://docs.imgproxy.net/)                     | /images/       |
| ImageProxy.Key            | IMAGE_PROXY_KEY             | [Signing key for Imgproxy](https://docs.imgproxy.net/configuration?id=url-signature) |                |
| ImageProxy.Salt           | IMAGE_PROXY_SALT            | [Salt for Imgproxy](https://docs.imgproxy.net/configuration?id=url-signature)        |                |
| ImageProxy.AllowedConfigs | IMAGE_PROXY_ALLOWED_CONFIGS | Comma-separated list of allowed image configs in format width:height:resizingType.   |                |
|                           |                             | Example `100:100:FILL,300:200:FIT`. Use `*` if not limiting the configs.             |                |
| Server.Debug              | GOCOM_DEBUG                 | Set to true to enable debug logging and put API framework in debug mode.             | false          |
| Server.Port               | GOCOM_PORT                  | HTTP port on which the GoCommerce API listens                                        | 8080           |
| Storage.Images.Type       | STORAGE_IMAGES_TYPE         | Type of storage used for storing images. Currently only `fs` is supported.           | fs             |
| Storage.Images.Path       | STORAGE_IMAGES_PATH         | Path for storing images                                                              | ./files/images |

## Running locally

```bash
./dc-up.sh
cd api
go run .
./dc-down.sh
```
