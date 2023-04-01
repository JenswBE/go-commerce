package fixtures

import (
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
