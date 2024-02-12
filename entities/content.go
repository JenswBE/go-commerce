package entities

import (
	"github.com/JenswBE/go-commerce/api/openapi"
)

// Content data
type Content struct {
	Name        string
	ContentType ContentType
	Body        string
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

type ContentType string

const (
	ContentTypeSimple = "SIMPLE"
	ContentTypeHTML   = "HTML"
)

func (contentType ContentType) String() string {
	return string(contentType)
}

func (contentType ContentType) IsValid() bool {
	return map[ContentType]bool{
		ContentTypeSimple: true,
		ContentTypeHTML:   true,
	}[contentType]
}
