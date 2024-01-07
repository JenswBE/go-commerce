package product

import (
	"github.com/jinzhu/copier"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
)

func ServiceCategoryFromEntity(p *presenter.Presenter, input *entities.ServiceCategory) openapi.ServiceCategory {
	output := openapi.NewServiceCategory(p.EncodeID(input.ID), p.String(input.Name), int64(input.Order))
	return *output
}

func ServiceFromEntity(p *presenter.Presenter, input *entities.Service) openapi.Service {
	output := openapi.NewService(
		p.EncodeID(input.ID),
		p.String(input.Name),
		p.String(input.Description),
		int64(input.Price.Int()),
		int64(input.Order),
	)
	return *output
}

func ResolvedServiceCategoryFromEntity(p *presenter.Presenter, input *entities.ResolvedServiceCategory) (openapi.ResolvedServiceCategory, error) {
	// Convert to basic service category
	serviceCategory := ServiceCategoryFromEntity(p, &input.ServiceCategory)
	output := openapi.ResolvedServiceCategory{}
	err := copier.Copy(&output, &serviceCategory)
	if err != nil {
		return openapi.ResolvedServiceCategory{}, entities.NewError(500, openapi.GOCOMERRORCODE_UNKNOWN_ERROR, input.ID.String(), err)
	}

	// Set services
	output.Services = presenter.SliceFromEntity(p, input.Services, ServiceFromEntity)
	return output, nil
}

func ResolvedServiceCategoryListFromEntity(p *presenter.Presenter, input []*entities.ResolvedServiceCategory) (openapi.ResolvedServiceCategoryList, error) {
	output := make([]openapi.ResolvedServiceCategory, len(input))
	var err error
	for i, svcCat := range input {
		output[i], err = ResolvedServiceCategoryFromEntity(p, svcCat)
		if err != nil {
			return openapi.ResolvedServiceCategoryList{}, err
		}
	}
	return *openapi.NewResolvedServiceCategoryList(output), nil
}
