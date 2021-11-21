package content

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/rs/zerolog/log"
)

const defaultContentType = openapi.CONTENTTYPE_SIMPLE

func ContentFromEntity(p *presenter.Presenter, input *entities.Content) openapi.Content {
	// Convert ContentType
	contentType, err := openapi.NewContentTypeFromValue(input.ContentType.String())
	if err != nil {
		log.Warn().Err(err).Stringer("content_type", input.ContentType).Msgf("Unknown content type received from entity, defaulting to %s", defaultContentType)
		contentType = defaultContentType.Ptr()
	}

	// Build OpenAPI model
	return *openapi.NewContent(input.Name, *contentType, input.Body)
}

func ContentSliceFromEntity(p *presenter.Presenter, input []*entities.Content) []openapi.Content {
	output := make([]openapi.Content, 0, len(input))
	for _, content := range input {
		output = append(output, ContentFromEntity(p, content))
	}
	return output
}

func ContentListFromEntity(p *presenter.Presenter, input []*entities.Content) openapi.ContentList {
	return *openapi.NewContentList(ContentSliceFromEntity(p, input))
}

func ContentToEntity(p *presenter.Presenter, content openapi.Content) *entities.Content {
	// Build entity
	return &entities.Content{
		Name:        content.Name,
		ContentType: entities.ContentType(content.GetContentType()),
		Body:        content.Body,
	}
}
