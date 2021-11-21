package product

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

func ManufacturerFromEntity(p *presenter.Presenter, input *entities.Manufacturer) openapi.Manufacturer {
	// Set basic fields
	output := openapi.NewManufacturer(input.Name)
	output.SetId(p.EncodeID(input.ID))
	output.SetWebsiteUrl(input.WebsiteURL)

	// Set image URL
	if input.Image != nil && len(input.Image.URLs) > 0 {
		output.SetImageUrls(input.Image.URLs)
	}
	return *output
}

func ManufacturerSliceFromEntity(p *presenter.Presenter, input []*entities.Manufacturer) []openapi.Manufacturer {
	output := make([]openapi.Manufacturer, 0, len(input))
	for _, manufacturer := range input {
		output = append(output, ManufacturerFromEntity(p, manufacturer))
	}
	return output
}

func ManufacturerListFromEntity(p *presenter.Presenter, input []*entities.Manufacturer) openapi.ManufacturerList {
	return *openapi.NewManufacturerList(ManufacturerSliceFromEntity(p, input))
}

func ManufacturerToEntity(p *presenter.Presenter, id uuid.UUID, input openapi.Manufacturer) *entities.Manufacturer {
	// Build entity
	return &entities.Manufacturer{
		ID:         id,
		Name:       input.GetName(),
		WebsiteURL: input.GetWebsiteUrl(),
	}
}
