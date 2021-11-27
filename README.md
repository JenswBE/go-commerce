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

| Config key                            | Env variable                    | Description                                                                                                | Default value  |
| ------------------------------------- | ------------------------------- | ---------------------------------------------------------------------------------------------------------- | -------------- |
| Authentication.IssuerURL              | AUTH_ISSUER_URL                 | URL to OpenID Configuration Issuer (without `.well-known/openid-configuration`)                            |                |
| Database.Default.Host                 | DATABASE_DEFAULT_HOST           | Hostname of the default Postgres datatabase                                                                |                |
| Database.Default.Port                 | DATABASE_DEFAULT_PORT           | Port of the default Postgres datatabase                                                                    | 5432           |
| Database.Default.User                 | DATABASE_DEFAULT_USER           | Username for the default Postgres datatabase                                                               |                |
| Database.Default.Password             | DATABASE_DEFAULT_PASSWORD       | Password for the default Postgres datatabase                                                               |                |
| Database.Default.Database             | DATABASE_DEFAULT_DATABASE       | Database name for default the Postgres datatabase                                                          |                |
| Database.Content.Host                 | DATABASE_CONTENT_HOST           | Override the default hostname for the content Postgres datatabase                                          |                |
| Database.Content.Port                 | DATABASE_CONTENT_PORT           | Override the default port for the content Postgres datatabase                                              |                |
| Database.Content.User                 | DATABASE_CONTENT_USER           | Override the default user for the content Postgres datatabase                                              |                |
| Database.Content.Password             | DATABASE_CONTENT_PASSWORD       | Override the default password for the content Postgres datatabase                                          |                |
| Database.Content.Database             | DATABASE_CONTENT_DATABASE       | Override the default database for the content Postgres datatabase                                          |                |
| Database.Product.Host                 | DATABASE_PRODUCT_HOST           | Override the default hostname for the product Postgres datatabase                                          |                |
| Database.Product.Port                 | DATABASE_PRODUCT_PORT           | Override the default port for the product Postgres datatabase                                              |                |
| Database.Product.User                 | DATABASE_PRODUCT_USER           | Override the default user for the product Postgres datatabase                                              |                |
| Database.Product.Password             | DATABASE_PRODUCT_PASSWORD       | Override the default password for the product Postgres datatabase                                          |                |
| Database.Product.Database             | DATABASE_PRODUCT_DATABASE       | Override the default database for the product Postgres datatabase                                          |                |
| Features.Categories.Enabled           | FEATURES_CATEGORIES_ENABLED     | Support for categories is enabled                                                                          | true           |
| Features.Manufacturers.Enabled        | FEATURES_MANUFACTURERS_ENABLED  | Support for manufacturers is enabled                                                                       | true           |
| Features.Products.Enabled             | FEATURES_PRODUCTS_ENABLED       | Support for products is enabled                                                                            | true           |
| Features.Content.Enabled              | FEATURES_CONTENT_ENABLED        | Support for content is enabled                                                                             | true           |
| Features.Content.List                 | FEATURES_CONTENT_LIST           | List of content. New content is automatically added to the DB. Missing content is not removed from the DB. |                |
|                                       |                                 | Config: Object with fields `Name` and `ContentType`                                                        |                |
|                                       |                                 | Env: List of format `Name:ContentType`                                                                     |                |
| Features.Events.Enabled               | FEATURES_EVENTS_ENABLED         | Support for events is enabled                                                                              | true           |
| Features.Content.Events.WholeDaysOnly | FEATURES_EVENTS_WHOLE_DAYS_ONLY | Only events with full days (no time) are supported                                                         | true           |
| ImageProxy.BaseURL                    | IMAGE_PROXY_BASE_URL            | Base URL of your [Imgproxy instance](https://docs.imgproxy.net/)                                           | /images/       |
| ImageProxy.Key                        | IMAGE_PROXY_KEY                 | [Signing key for Imgproxy](https://docs.imgproxy.net/configuration?id=url-signature)                       |                |
| ImageProxy.Salt                       | IMAGE_PROXY_SALT                | [Salt for Imgproxy](https://docs.imgproxy.net/configuration?id=url-signature)                              |                |
| ImageProxy.AllowedConfigs             | IMAGE_PROXY_ALLOWED_CONFIGS     | Comma-separated list of allowed image configs in format width:height:resizingType.                         |                |
|                                       |                                 | Example `100:100:FILL,300:200:FIT`. Use `*` if not limiting the configs.                                   |                |
| Server.Debug                          | GOCOM_DEBUG                     | Set to true to enable debug logging and put API framework in debug mode.                                   | false          |
| Server.Port                           | GOCOM_PORT                      | HTTP port on which the GoCommerce API listens                                                              | 8080           |
| Storage.Images.Type                   | STORAGE_IMAGES_TYPE             | Type of storage used for storing images. Currently only `fs` is supported.                                 | fs             |
| Storage.Images.Path                   | STORAGE_IMAGES_PATH             | Path for storing images                                                                                    | ./files/images |

## Running locally

```bash
docker compose up -d
cd api
go run .
docker compose down
```
