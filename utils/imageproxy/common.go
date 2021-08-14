package imageproxy

import (
	"fmt"
	"strconv"
)

type ImageConfig struct {
	Width        int
	Height       int
	ResizingType ResizingType
}

func ParseImageConfig(width, height, resizingType string) (ImageConfig, error) {
	// Init vars
	config := ImageConfig{}
	var err error

	// Parse width
	config.Width, err = strconv.Atoi(width)
	if err != nil {
		return ImageConfig{}, fmt.Errorf(`failed to parse width as integer: %s`, width)
	}

	// Parse height
	config.Height, err = strconv.Atoi(height)
	if err != nil {
		return ImageConfig{}, fmt.Errorf(`failed to parse height as integer: %s`, height)
	}

	// Parse resizing type
	config.ResizingType, err = ParseResizingType(resizingType)
	if err != nil {
		return ImageConfig{}, err
	}

	// Validate config
	err = config.Validate()
	if err != nil {
		return ImageConfig{}, err
	}

	// Parse successful
	return config, nil
}

func (config ImageConfig) Validate() error {
	if config.Width <= 0 {
		return fmt.Errorf(`width cannot be negative or zero: %d`, config.Width)
	}
	if config.Height <= 0 {
		return fmt.Errorf(`height cannot be negative or zero: %d`, config.Height)
	}
	return nil
}

func (config ImageConfig) String() string {
	return fmt.Sprintf(`%d:%d:%s`, config.Width, config.Height, config.ResizingType)
}
