package fixtures

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/config"
)

// #############################
// #           ENTITY          #
// #############################

func Config() config.Config {
	fix := config.Config{}
	fix.Features.Categories.Enabled = true
	fix.Features.Manufacturers.Enabled = true
	fix.Features.Products.Enabled = true
	fix.Features.Content.Enabled = true
	fix.Features.Events.Enabled = true
	fix.Features.Events.WholeDaysOnly = true
	return fix
}

// #############################
// #          OPENAPI          #
// #############################

func ConfigOpenAPI() openapi.Config {
	return openapi.Config{
		Features: openapi.ConfigFeatures{
			Categories: openapi.ConfigFeaturesCategories{
				Enabled: true,
			},
			Manufacturers: openapi.ConfigFeaturesManufacturers{
				Enabled: true,
			},
			Products: openapi.ConfigFeaturesProducts{
				Enabled: true,
			},
			Content: openapi.ConfigFeaturesContent{
				Enabled: true,
			},
			Events: openapi.ConfigFeaturesEvents{
				Enabled:       true,
				WholeDaysOnly: true,
			},
		},
	}
}
