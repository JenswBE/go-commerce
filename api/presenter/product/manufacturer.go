package product

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
)

func ManufacturerFromEntity(p *presenter.Presenter, input *entities.Manufacturer) openapi.Manufacturer {
	// Set basic fields
	output := openapi.NewManufacturer(p.EncodeID(input.ID), p.String(input.Name))
	output.SetWebsiteUrl(p.String(input.WebsiteURL))

	// Set image URL
	if input.Image != nil && len(input.Image.URLs) > 0 {
		output.SetImageUrls(input.Image.URLs)
	}
	return *output
}

func ManufacturerListFromEntity(p *presenter.Presenter, input []*entities.Manufacturer) openapi.ManufacturerList {
	return *openapi.NewManufacturerList(presenter.SliceFromEntity(p, input, ManufacturerFromEntity))
}
