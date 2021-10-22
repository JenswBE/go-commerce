package entities

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
)

type Image struct {
	ID        ID
	Extension string // File extension
	URLs      map[string]string
	Order     int
}

// Validate validates the product data
func (img *Image) Validate() error {
	// Validate simple fields
	if img.Order < 0 {
		return NewError(400, openapi.ERRORCODE_IMAGE_ORDER_NEGATIVE, img.ID.String(), nil)
	}

	// Entity is valid
	return nil
}

// SetURLsFromConfigs generates and sets an URL from the provided config
func (img *Image) SetURLsFromConfigs(service imageproxy.Service, configs map[string]imageproxy.ImageConfig) error {
	// Skip if img is nil
	if img == nil {
		return nil
	}

	// Generate URL's
	img.URLs = make(map[string]string, len(configs))
	localSource := "local:///" + img.ID.String() + img.Extension
	var err error
	for request, config := range configs {
		img.URLs[request], err = service.GenerateURL(localSource, config)
		if err != nil {
			return err
		}
	}
	return nil
}
