package entities

import "github.com/JenswBE/go-commerce/api/openapi"

// Content data
type Content struct {
	Name        string
	ContentType ContentType
	Content     string
}

type ContentType string

const ContentTypeSimple = "SIMPLE"
const ContentTypeMarkdown = "MARKDOWN"

func (contenType ContentType) String() string {
	return string(contenType)
}

// Validate validates the content data
func (c *Content) Validate() error {
	// Validate simple fields
	if c.Name == "" {
		return NewError(400, openapi.GOCOMERRORCODE_CONTENT_NAME_EMPTY, c.Name, nil)
	}

	// Entity is valid
	return nil
}
