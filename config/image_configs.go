package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/JenswBE/go-commerce/utils/imageproxy"
)

func ParseAllowedImageConfigs(configs string) ([]imageproxy.ImageConfig, error) {
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
