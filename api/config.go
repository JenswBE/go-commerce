package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/go-playground/validator/v10"
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
	ImageProxy struct {
		BaseURL string
		Key     string
		Salt    string

		// Comma-separated list of allowed configs in format width:height:resizingType.
		// Example "100:100:FILL,300:200:FIT". Use "*" if not limiting the configs.
		AllowedConfigs string
	}
	Server struct {
		Debug bool
		Port  int
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

func parseConfig() (*Config, error) {
	// Set defaults
	viper.SetDefault("Database.Default.Port", 5432)
	viper.SetDefault("ImageProxy.BaseURL", "/images/")
	viper.SetDefault("Server.Port", 8080)
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
	viper.BindEnv("Authentication.IssuerURL", "AUTH_ISSUER_URL")
	viper.BindEnv("Database.Default.Host", "DATABASE_DEFAULT_HOST")
	viper.BindEnv("Database.Default.Port", "DATABASE_DEFAULT_PORT")
	viper.BindEnv("Database.Default.User", "DATABASE_DEFAULT_USER")
	viper.BindEnv("Database.Default.Password", "DATABASE_DEFAULT_PASSWORD")
	viper.BindEnv("Database.Default.Database", "DATABASE_DEFAULT_DATABASE")
	viper.BindEnv("Database.Content.Host", "DATABASE_CONTENT_HOST")
	viper.BindEnv("Database.Content.Port", "DATABASE_CONTENT_PORT")
	viper.BindEnv("Database.Content.User", "DATABASE_CONTENT_USER")
	viper.BindEnv("Database.Content.Password", "DATABASE_CONTENT_PASSWORD")
	viper.BindEnv("Database.Content.Database", "DATABASE_CONTENT_DATABASE")
	viper.BindEnv("Database.Product.Host", "DATABASE_PRODUCT_HOST")
	viper.BindEnv("Database.Product.Port", "DATABASE_PRODUCT_PORT")
	viper.BindEnv("Database.Product.User", "DATABASE_PRODUCT_USER")
	viper.BindEnv("Database.Product.Password", "DATABASE_PRODUCT_PASSWORD")
	viper.BindEnv("Database.Product.Database", "DATABASE_PRODUCT_DATABASE")
	viper.BindEnv("ImageProxy.BaseURL", "IMAGE_PROXY_BASE_URL")
	viper.BindEnv("ImageProxy.Key", "IMAGE_PROXY_KEY")
	viper.BindEnv("ImageProxy.Salt", "IMAGE_PROXY_SALT")
	viper.BindEnv("ImageProxy.AllowedConfigs", "IMAGE_PROXY_ALLOWED_CONFIGS")
	viper.BindEnv("Server.Debug", "GOCOM_DEBUG")
	viper.BindEnv("Server.Port", "GOCOM_PORT")
	viper.BindEnv("Storage.Images.Type", "STORAGE_IMAGES_TYPE")
	viper.BindEnv("Storage.Images.Path", "STORAGE_IMAGES_PATH")

	// Unmarshal to Config struct
	var config Config
	err = viper.Unmarshal(&config)
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

func buildDSN(config Database, fallback Database) string {
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

func parseAllowedImageConfigs(configs string) ([]imageproxy.ImageConfig, error) {
	// Configs cannot be empty, wildcard should be used in this case
	if configs == "" {
		return nil, errors.New(`allowed image configs cannot be empty, use wildcard * instead`)
	}

	// Return empty list on wildcard
	if configs == "*" {
		return []imageproxy.ImageConfig{}, nil
	}

	// Split config string in chunks
	configChunks := strings.Split(configs, ",")

	// Build image configs
	imgConfigs := make([]imageproxy.ImageConfig, 0, len(configChunks))
	for _, chunk := range configChunks {
		// Split chunk in parts
		parts := strings.Split(chunk, ":")

		// Each chunk must consist of 3 parts (width:height:resizingType)
		if len(parts) != 3 {
			return nil, fmt.Errorf(`chunk should consist of 3 parts width:height:resizingType, received %s`, chunk)
		}

		// Parse config
		config, err := imageproxy.ParseImageConfig(parts[0], parts[1], parts[2])
		if err != nil {
			return nil, err
		}
		imgConfigs = append(imgConfigs, config)
	}

	// Parse successful
	return imgConfigs, nil
}
