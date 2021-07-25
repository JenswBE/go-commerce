package entities

import (
	"errors"

	"github.com/JenswBE/go-commerce/utils/imageproxy"
)

type Image struct {
	ID        ID
	Extension string // File extension
	URL       string
	Order     int
}

// Validate validates the product data
func (img *Image) Validate() error {
	// Validate simple fields
	if img.Order < 0 {
		return NewError(400, errors.New("image order cannot be negative"))
	}

	// Entity is valid
	return nil
}

// SetURLFromConfig generates and sets an URL from the provided config
func (img *Image) SetURLFromConfig(service imageproxy.Service, config imageproxy.ImageConfig) error {
	// Generate URL
	localSource := "local:///" + img.ID.String() + img.Extension
	url, err := service.GenerateURL(localSource, config)
	if err != nil {
		return err
	}

	// Set URL
	img.URL = url
	return nil
}
