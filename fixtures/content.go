package fixtures

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
)

// #############################
// #           ENTITY          #
// #############################

func Content() *entities.Content {
	return &entities.Content{
		Name:        "test-name",
		ContentType: entities.ContentTypeHTML,
		Body:        "test-body",
	}
}

func ContentSlice() []*entities.Content {
	return []*entities.Content{
		Content(),
	}
}

// #############################
// #          OPENAPI          #
// #############################

func ContentOpenAPI() *openapi.Content {
	return &openapi.Content{
		Name:        "test-name",
		ContentType: openapi.CONTENTTYPE_HTML,
		Body:        "test-body",
	}
}

func ContentOpenAPISlice() []openapi.Content {
	return []openapi.Content{
		*ContentOpenAPI(),
	}
}

func ContentListOpenAPI() *openapi.ContentList {
	return openapi.NewContentList(ContentOpenAPISlice())
}
