package product

import (
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
)

func ImageURLsSliceFromEntity(p *presenter.Presenter, input []*entities.Image) []map[string]string {
	output := make([]map[string]string, 0, len(input))
	for _, image := range input {
		output = append(output, image.URLs)
	}
	return output
}
