package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
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
		Port int
	}
	Storage struct {
		Images Storage
	}
}

type Storage struct {
	Type string
	Path string
}

const StorageTypeFS = "fs"

func parseConfig() (*Config, error) {
	// Set defaults
	viper.SetDefault("Database.Port", 5432)
	viper.SetDefault("ImageProxy.BaseURL", "/images/")
	viper.SetDefault("Server.Port", 8090)
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
	viper.BindEnv("Database.Host", "DATABASE_HOST")
	viper.BindEnv("Database.Port", "DATABASE_PORT")
	viper.BindEnv("Database.User", "DATABASE_USER")
	viper.BindEnv("Database.Password", "DATABASE_PASSWORD")
	viper.BindEnv("Database.Database", "DATABASE_DATABASE")
	viper.BindEnv("Server.Port", "GOCOM_PORT")
	viper.BindEnv("Storage.Images.Type", "STORAGE_IMAGES_TYPE")
	viper.BindEnv("Storage.Images.Path", "STORAGE_IMAGES_PATH")

	// Unmarshal to Config struct
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}
	return &config, nil
}

func buildDSN(config *Config) string {
	options := make([]string, 0, 5)
	if config.Database.Host != "" {
		options = append(options, "host="+config.Database.Host)
	}
	if config.Database.Port > 0 {
		options = append(options, "port="+strconv.Itoa(config.Database.Port))
	}
	if config.Database.User != "" {
		options = append(options, "user="+config.Database.User)
	}
	if config.Database.Password != "" {
		options = append(options, "password="+config.Database.Password)
	}
	if config.Database.Database != "" {
		options = append(options, "dbname="+config.Database.Database)
	}
	return strings.Join(options, " ")
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
