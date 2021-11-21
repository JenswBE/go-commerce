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
		ContentType: entities.ContentTypeMarkdown,
		Content:     "test-content",
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
		ContentType: openapi.CONTENTTYPE_MARKDOWN,
		Content:     "test-content",
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
