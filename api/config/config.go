package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Authentication struct {
		IssuerURL string `validate:"required"`
	}
	Database struct {
		Default Database
		Content Database
		Product Database
	}
	Features struct {
		Categories struct {
			Enabled bool
		}
		Manufacturers struct {
			Enabled bool
		}
		Products struct {
			Enabled bool
		}
		Content struct {
			Enabled bool
			List    ContentList
		}
		Events struct {
			Enabled       bool
			WholeDaysOnly bool
		}
	}
	ImageProxy struct {
		BaseURL string
		Key     string
		Salt    string

		// Comma-separated list of allowed configs in format width:height:resizingType.
		// Example "100:100:FILL,300:200:FIT". Use "*" if not limiting the configs.
		AllowedConfigs string
	}
	Server struct {
		Debug          bool
		Port           int
		TrustedProxies []string
	}
	Storage struct {
		Images Storage
	}
}

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Storage struct {
	Type string
	Path string
}

const StorageTypeFS = "fs"

func ParseConfig() (*Config, error) {
	// Set defaults
	viper.SetDefault("Database.Default.Port", 5432)
	viper.SetDefault("Features.Categories.Enabled", true)
	viper.SetDefault("Features.Manufacturers.Enabled", true)
	viper.SetDefault("Features.Products.Enabled", true)
	viper.SetDefault("Features.Content.Enabled", true)
	viper.SetDefault("Features.Events.Enabled", true)
	viper.SetDefault("ImageProxy.BaseURL", "/images/")
	viper.SetDefault("Server.Port", 8080)
	viper.SetDefault("Server.TrustedProxies", []string{"172.16.0.0/16"}) // Default Docker IP range
	viper.SetDefault("Storage.Images.Type", StorageTypeFS)
	viper.SetDefault("Storage.Images.Path", "./files/images")

	// Parse config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed reading config file: %w", err)
		}
		log.Warn().Err(err).Msg("No config file found, expecting configuration through ENV variables")
	}

	// Bind ENV variables
	err = bindEnvs([]envBinding{
		{"Authentication.IssuerURL", "AUTH_ISSUER_URL"},
		{"Database.Default.Host", "DATABASE_DEFAULT_HOST"},
		{"Database.Default.Port", "DATABASE_DEFAULT_PORT"},
		{"Database.Default.User", "DATABASE_DEFAULT_USER"},
		{"Database.Default.Password", "DATABASE_DEFAULT_PASSWORD"},
		{"Database.Default.Database", "DATABASE_DEFAULT_DATABASE"},
		{"Database.Content.Host", "DATABASE_CONTENT_HOST"},
		{"Database.Content.Port", "DATABASE_CONTENT_PORT"},
		{"Database.Content.User", "DATABASE_CONTENT_USER"},
		{"Database.Content.Password", "DATABASE_CONTENT_PASSWORD"},
		{"Database.Content.Database", "DATABASE_CONTENT_DATABASE"},
		{"Database.Product.Host", "DATABASE_PRODUCT_HOST"},
		{"Database.Product.Port", "DATABASE_PRODUCT_PORT"},
		{"Database.Product.User", "DATABASE_PRODUCT_USER"},
		{"Database.Product.Password", "DATABASE_PRODUCT_PASSWORD"},
		{"Database.Product.Database", "DATABASE_PRODUCT_DATABASE"},
		{"Features.Categories.Enabled", "FEATURES_CATEGORIES_ENABLED"},
		{"Features.Manufacturers.Enabled", "FEATURES_MANUFACTURERS_ENABLED"},
		{"Features.Products.Enabled", "FEATURES_PRODUCTS_ENABLED"},
		{"Features.Content.Enabled", "FEATURES_CONTENT_ENABLED"},
		{"Features.Content.List", "FEATURES_CONTENT_LIST"},
		{"Features.Events.Enabled", "FEATURES_EVENTS_ENABLED"},
		{"Features.Content.Events.WholeDaysOnly", "FEATURES_EVENTS_WHOLE_DAYS_ONLY"},
		{"ImageProxy.BaseURL", "IMAGE_PROXY_BASE_URL"},
		{"ImageProxy.Key", "IMAGE_PROXY_KEY"},
		{"ImageProxy.Salt", "IMAGE_PROXY_SALT"},
		{"ImageProxy.AllowedConfigs", "IMAGE_PROXY_ALLOWED_CONFIGS"},
		{"Server.Debug", "GOCOM_DEBUG"},
		{"Server.Port", "GOCOM_PORT"},
		{"Server.TrustedProxies", "GOCOM_TRUSTED_PROXIES"},
		{"Storage.Images.Type", "STORAGE_IMAGES_TYPE"},
		{"Storage.Images.Path", "STORAGE_IMAGES_PATH"},
	})
	if err != nil {
		return nil, err
	}

	// Unmarshal to Config struct
	var config Config
	decodeHooks := mapstructure.ComposeDecodeHookFunc(
		contentListHook(),
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
	)
	err = viper.Unmarshal(&config, viper.DecodeHook(decodeHooks))
	if err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	// Validate config
	validate := validator.New()
	if err := validate.Struct(&config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}
	return &config, nil
}

type envBinding struct {
	configPath string
	envName    string
}

func bindEnvs(bindings []envBinding) error {
	for _, binding := range bindings {
		err := viper.BindEnv(binding.configPath, binding.envName)
		if err != nil {
			return fmt.Errorf("failed to bind env var %s to %s: %w", binding.envName, binding.configPath, err)
		}
	}
	return nil
}

func BuildDSN(config Database, fallback Database) string {
	options := make([]string, 0, 5)
	host := stringFallback(config.Host, fallback.Host)
	if host != "" {
		options = append(options, "host="+host)
	}
	port := intFallback(config.Port, fallback.Port)
	if port > 0 {
		options = append(options, "port="+strconv.Itoa(port))
	}
	user := stringFallback(config.User, fallback.User)
	if user != "" {
		options = append(options, "user="+user)
	}
	password := stringFallback(config.Password, fallback.Password)
	if password != "" {
		options = append(options, "password="+password)
	}
	database := stringFallback(config.Database, fallback.Database)
	if database != "" {
		options = append(options, "dbname="+database)
	}
	return strings.Join(options, " ")
}

func stringFallback(value, fallback string) string {
	if value != "" {
		return value
	}
	return fallback
}

func intFallback(value, fallback int) int {
	if value != 0 {
		return value
	}
	return fallback
}
