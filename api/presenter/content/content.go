package content

import (
	"github.com/rs/zerolog/log"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
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
	var body string
	switch *contentType {
	case openapi.CONTENTTYPE_SIMPLE:
		body = p.String(input.Body)
	case openapi.CONTENTTYPE_HTML:
		body = p.ContentHTML(input.Body)
	}
	return *openapi.NewContent(p.String(input.Name), *contentType, body)
}

func ContentSliceFromEntity(p *presenter.Presenter, input []*entities.Content) []openapi.Content {
	output := make([]openapi.Content, 0, len(input))
	for _, content := range input {
		output = append(output, ContentFromEntity(p, content))
	}
	return output
}
