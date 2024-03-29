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

func ContentEntityToPg(c *entities.Content) *Content {
	return &Content{
		Name:        c.Name,
		ContentType: c.ContentType.String(),
		Body:        c.Body,
	}
}
