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
}

func parseConfig() (*Config, error) {
	// Set defaults
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("server.port", 8090)

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
	viper.BindEnv("database.host", "DATABASE_HOST")
	viper.BindEnv("database.port", "DATABASE_PORT")
	viper.BindEnv("database.user", "DATABASE_USER")
	viper.BindEnv("database.password", "DATABASE_PASSWORD")
	viper.BindEnv("database.database", "DATABASE_DATABASE")
	viper.BindEnv("server.port", "GOCOM_PORT")

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
