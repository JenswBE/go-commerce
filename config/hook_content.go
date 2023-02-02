package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"

	"github.com/JenswBE/go-commerce/entities"
)

type Content struct {
	Name        string
	ContentType string
}

type ContentList []Content

func contentListHook() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type,
		t reflect.Type,
		data any,
	) (any, error) {
		// Check that the data is string
		if f.Kind() != reflect.String {
			return data, nil
		}

		// Check that the target type is our custom type
		if t != reflect.TypeOf(ContentList{}) {
			return data, nil
		}

		// Split content string in chunks.
		// Type of "data" is already checked.
		content, _ := data.(string)
		contentChunks := strings.Split(content, ",")

		// Build content list
		contentList := make([]Content, 0, len(contentChunks))
		for _, chunk := range contentChunks {
			// Split chunk in parts
			parts := strings.Split(chunk, ":")

			// Each chunk must consist of 2 parts (name:contentType)
			if len(parts) != 2 {
				return nil, fmt.Errorf(`chunk should consist of 2 parts name:contentType, received %s`, chunk)
			}

			// Parse content
			contentList = append(contentList, Content{
				Name:        parts[0],
				ContentType: parts[1],
			})
		}

		// Parse successful
		return contentList, nil
	}
}

func (c Content) ToEntity() entities.Content {
	return entities.Content{
		Name:        c.Name,
		ContentType: entities.ContentType(c.ContentType),
		Body:        "",
	}
}

func (cl ContentList) ToEntity() []entities.Content {
	output := make([]entities.Content, 0, len(cl))
	for _, content := range cl {
		output = append(output, content.ToEntity())
	}
	return output
}
