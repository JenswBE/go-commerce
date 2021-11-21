package internal

import (
	"github.com/JenswBE/go-commerce/entities"
)

type Content struct {
	Name        string `gorm:"primaryKey"`
	ContentType string
	Body        string
}

func (c *Content) ToEntity() *entities.Content {
	return &entities.Content{
		Name:        c.Name,
		ContentType: entities.ContentType(c.ContentType),
		Body:        c.Body,
	}
}

func ContentsListPgToEntity(c []*Content) []*entities.Content {
	output := make([]*entities.Content, 0, len(c))
	for _, content := range c {
		output = append(output, content.ToEntity())
	}
	return output
}

func ContentEntityToPg(c *entities.Content) *Content {
	return &Content{
		Name:        c.Name,
		ContentType: c.ContentType.String(),
		Body:        c.Body,
	}
}
