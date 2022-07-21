package content

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/config"
)

func ConfigFromEntity(p *presenter.Presenter, input config.Config) openapi.Config {
	configFeaturesCategories := openapi.NewConfigFeaturesCategories(input.Features.Categories.Enabled)
	configFeaturesManufacturers := openapi.NewConfigFeaturesManufacturers(input.Features.Manufacturers.Enabled)
	configFeaturesProducts := openapi.NewConfigFeaturesProducts(input.Features.Products.Enabled)
	configFeaturesContent := openapi.NewConfigFeaturesContent(input.Features.Content.Enabled)
	configFeaturesEvents := openapi.NewConfigFeaturesEvents(input.Features.Events.Enabled, input.Features.Events.WholeDaysOnly)
	configFeatures := openapi.NewConfigFeatures(
		*configFeaturesCategories,
		*configFeaturesManufacturers,
		*configFeaturesProducts,
		*configFeaturesContent,
		*configFeaturesEvents,
	)
	return *openapi.NewConfig(*configFeatures)
}
