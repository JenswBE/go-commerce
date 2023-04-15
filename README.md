[![GitHub](https://img.shields.io/github/license/JenswBE/go-commerce)](https://github.com/JenswBE/go-commerce)
[![Docker Pulls](https://img.shields.io/docker/pulls/jenswbe/go-commerce)](https://hub.docker.com/r/jenswbe/go-commerce)

# GoCommerce

KISS backend for an e-commerce

## Links

- GitHub: https://github.com/JenswBE/go-commerce
- DockerHub: https://hub.docker.com/r/jenswbe/go-commerce

## Configuration

GoCommerce can be configured in 2 ways:

1. Create a file called `config.yml` in the same folder or the parent folder of the binary. See `config.yml` for an example.
2. Set environment variables

If both are defined, the environment variables take precedence.

| Config key                             | Env variable                             | Description                                                                                                | Default value  |
| -------------------------------------- | ---------------------------------------- | ---------------------------------------------------------------------------------------------------------- | -------------- |
| Authentication.Type                    | AUTH_TYPE                                | Switch between OIDC and NONE (latter should only be used for testing)                                      | OIDC           |
| Authentication.OIDC.IssuerURL          | AUTH_OIDC_ISSUER_URL                     | URL to OpenID Configuration Issuer (without `.well-known/openid-configuration`)                            |                |
| Authentication.OIDC.ClientID           | AUTH_OIDC_CLIENT_ID                      | Client ID for OIDC                                                                                         |                |
| Authentication.OIDC.ClientSecret       | AUTH_OIDC_CLIENT_SECRET                  | Client secret for OIDC                                                                                     |                |
| Authentication.SessionAuthKey          | AUTH_SESSION_AUTH_KEY                    | Authentication key for session tokens. Mandatory and must be a base64 encoded string of 64 bytes.          |                |
|                                        |                                          | Can be generated using `openssl rand -base64 64 \| paste --delimiters '' --serial`                         |                |
| Authentication.SessionEncKey           | AUTH_SESSION_ENC_KEY                     | Encryption key for session tokens. Mandatory and must be a base64 encoded string of 32 bytes.              |                |
|                                        |                                          | Can be generated using `openssl rand -base64 32 \| paste --delimiters '' --serial`                         |                |
| Database.Default.Host                  | DATABASE_DEFAULT_HOST                    | Hostname of the default Postgres datatabase                                                                |                |
| Database.Default.Port                  | DATABASE_DEFAULT_PORT                    | Port of the default Postgres datatabase                                                                    | 5432           |
| Database.Default.User                  | DATABASE_DEFAULT_USER                    | Username for the default Postgres datatabase                                                               |                |
| Database.Default.Password              | DATABASE_DEFAULT_PASSWORD                | Password for the default Postgres datatabase                                                               |                |
| Database.Default.Database              | DATABASE_DEFAULT_DATABASE                | Database name for default the Postgres datatabase                                                          |                |
| Database.Content.Host                  | DATABASE_CONTENT_HOST                    | Override the default hostname for the content Postgres datatabase                                          |                |
| Database.Content.Port                  | DATABASE_CONTENT_PORT                    | Override the default port for the content Postgres datatabase                                              |                |
| Database.Content.User                  | DATABASE_CONTENT_USER                    | Override the default user for the content Postgres datatabase                                              |                |
| Database.Content.Password              | DATABASE_CONTENT_PASSWORD                | Override the default password for the content Postgres datatabase                                          |                |
| Database.Content.Database              | DATABASE_CONTENT_DATABASE                | Override the default database for the content Postgres datatabase                                          |                |
| Database.Product.Host                  | DATABASE_PRODUCT_HOST                    | Override the default hostname for the product Postgres datatabase                                          |                |
| Database.Product.Port                  | DATABASE_PRODUCT_PORT                    | Override the default port for the product Postgres datatabase                                              |                |
| Database.Product.User                  | DATABASE_PRODUCT_USER                    | Override the default user for the product Postgres datatabase                                              |                |
| Database.Product.Password              | DATABASE_PRODUCT_PASSWORD                | Override the default password for the product Postgres datatabase                                          |                |
| Database.Product.Database              | DATABASE_PRODUCT_DATABASE                | Override the default database for the product Postgres datatabase                                          |                |
| Features.StartpageFeature              | FEATURES_STARTPAGE_FEATURE               | Feature which should be shown as startpage. See below config keys for supported features.                  | Products       |
| Features.Categories.Enabled            | FEATURES_CATEGORIES_ENABLED              | Support for categories is enabled                                                                          | true           |
| Features.Manufacturers.Enabled         | FEATURES_MANUFACTURERS_ENABLED           | Support for manufacturers is enabled                                                                       | true           |
| Features.Products.Enabled              | FEATURES_PRODUCTS_ENABLED                | Support for products is enabled                                                                            | true           |
| Features.Products.PublicURLTemplate    | FEATURES_PRODUCTS_PUBLIC_URL_TEMPLATE    | Optional template for showing link to product page on public site. Button is omitted in                    |                |
|                                        |                                          | list if not provided. String is parsed into a Go HTML template. Product is available as `.`.               |                |
| Features.Products.ShortDescriptionOnly | FEATURES_PRODUCTS_SHORT_DESCRIPTION_ONLY | Support for long descriptions for products is disabled                                                     | true           |
| Features.Content.Enabled               | FEATURES_CONTENT_ENABLED                 | Support for content is enabled                                                                             | true           |
| Features.Content.List                  | FEATURES_CONTENT_LIST                    | List of content. New content is automatically added to the DB. Missing content is not removed from the DB. |                |
|                                        |                                          | Config: Object with fields `Name` and `ContentType`                                                        |                |
|                                        |                                          | Env: List of format `Name:ContentType`                                                                     |                |
| Features.Events.Enabled                | FEATURES_EVENTS_ENABLED                  | Support for events is enabled                                                                              | true           |
| Features.Events.WholeDaysOnly          | FEATURES_EVENTS_WHOLE_DAYS_ONLY          | Only events with full days (no time) are supported                                                         | true           |
| ImageProxy.BaseURL                     | IMAGE_PROXY_BASE_URL                     | Base URL of your [Imgproxy instance](https://docs.imgproxy.net/)                                           | /images/       |
| ImageProxy.Key                         | IMAGE_PROXY_KEY                          | [Signing key for Imgproxy](https://docs.imgproxy.net/configuration?id=url-signature)                       |                |
| ImageProxy.Salt                        | IMAGE_PROXY_SALT                         | [Salt for Imgproxy](https://docs.imgproxy.net/configuration?id=url-signature)                              |                |
| ImageProxy.AllowedConfigs              | IMAGE_PROXY_ALLOWED_CONFIGS              | Comma-separated list of allowed image configs in format width:height:resizingType.                         |                |
|                                        |                                          | Example `100:100:FILL,300:200:FIT`. Use `*` if not limiting the configs.                                   |                |
| Server.Debug                           | GOCOM_DEBUG                              | Set to true to enable debug logging and put API framework in debug mode.                                   | false          |
| Server.Port                            | GOCOM_PORT                               | HTTP port on which the GoCommerce API listens                                                              | 8080           |
| Server.TrustedProxies                  | GOCOM_TRUSTED_PROXIES                    | IP's of proxies trusted by GoCommerce. Header `X-Forwarded-For` is only considered for these hosts.        | 172.16.0.0/16  |
| Storage.Images.Type                    | STORAGE_IMAGES_TYPE                      | Type of storage used for storing images. Currently only `fs` is supported.                                 | fs             |
| Storage.Images.Path                    | STORAGE_IMAGES_PATH                      | Path for storing images                                                                                    | ./files/images |

## Running locally

```bash
podman-compose up -d
go run .
podman-compose down
```

## Run end-to-end tests

```bash
# For Docker Compose use: "-f docker-compose.e2e.yml -f docker-compose.e2e.docker.yml"
podman-compose -f docker-compose.e2e.yml up -d
go test --tags e2e ./...
podman-compose -f docker-compose.e2e.yml down
```
