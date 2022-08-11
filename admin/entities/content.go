package entities

import (
	"html/template"

	"github.com/JenswBE/go-commerce/entities"
)

type ContentListData struct {
	BaseData
	Content []*entities.Content
}

type ContentFormData struct {
	BaseData
	ContentName   string
	IsHTMLContent bool
	Content       Content
}

type Content struct {
	BodySimple string        `form:"body"`
	BodyHTML   template.HTML `form:"body"`
}
