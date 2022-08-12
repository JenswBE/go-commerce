package entities

import (
	"html/template"

	"github.com/JenswBE/go-commerce/entities"
)

type ContentListTemplate struct {
	BaseData
	Content []*entities.Content
}

func (t ContentListTemplate) GetTemplateName() string {
	return "contentList"
}

type ContentFormTemplate struct {
	BaseData
	ContentName   string
	IsHTMLContent bool
	Content       Content
}

func (t ContentFormTemplate) GetTemplateName() string {
	return "contentForm"
}

type Content struct {
	BodySimple string        `form:"body"`
	BodyHTML   template.HTML `form:"body"`
}
