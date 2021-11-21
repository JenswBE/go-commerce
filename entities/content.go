package entities

import (
	"github.com/JenswBE/go-commerce/api/openapi"
)

// Content data
type Content struct {
	Name        string
	ContentType ContentType
	Content     string
}

type ContentType string

const ContentTypeSimple = "SIMPLE"
const ContentTypeMarkdown = "MARKDOWN"

func (contentType ContentType) String() string {
	return string(contentType)
}

func (contentType ContentType) IsValid() bool {
	// Valid types
	validContentTypes := []ContentType{
		ContentTypeSimple,
		ContentTypeMarkdown,
	}

	// Check provided type
	for _, validType := range validContentTypes {
		if validType == contentType {
			return true
		}
	}
	return false
}

// Validate validates the content data
func (c *Content) Validate() error {
	// Validate simple fields
	if c.Name == "" {
		return NewError(400, openapi.GOCOMERRORCODE_CONTENT_NAME_EMPTY, c.Name, nil)
	}
	if !c.ContentType.IsValid() {
		return NewError(400, openapi.GOCOMERRORCODE_CONTENT_TYPE_INVALID, c.Name, nil)
	}

	// Entity is valid
	return nil
}
