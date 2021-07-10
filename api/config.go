package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
	Server struct {
		Port int
	}
	Storage struct {
		Path         string
		ImagesFolder string
	}
}

func parseConfig() (*Config, error) {
	// Set defaults
	viper.SetDefault("Database.Port", 5432)
	viper.SetDefault("Server.Port", 8090)
	viper.SetDefault("Storage.Path", "./files")
	viper.SetDefault("Storage.ImagesFolder", "images")

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
		log.Printf("Warning: No config file found, expecting configuration through ENV variables %s", err)
	}

	// Bind ENV variables
	viper.BindEnv("Database.Host", "DATABASE_HOST")
	viper.BindEnv("Database.Port", "DATABASE_PORT")
	viper.BindEnv("Database.User", "DATABASE_USER")
	viper.BindEnv("Database.Password", "DATABASE_PASSWORD")
	viper.BindEnv("Database.Database", "DATABASE_DATABASE")
	viper.BindEnv("Server.Port", "GOCOM_PORT")
	viper.BindEnv("Storage.Path", "STORAGE_PATH")
	viper.BindEnv("Storage.ImagesFolder", "STORAGE_IMAGES_FOLDER")

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
