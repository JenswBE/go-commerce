package admin

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/JenswBE/go-commerce/admin/entities"
	baseEntities "github.com/JenswBE/go-commerce/entities"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const paramContentName = "content_name"

func (h *Handler) handleContentList(c *gin.Context) {
	// Fetch content
	content, err := h.contentService.ListContent()
	if err != nil {
		c.String(http.StatusInternalServerError, "Ophalen van inhoud mislukt: %v", err)
	}

	htmlWithFlashes(c, http.StatusOK, "contentList", &entities.ContentListData{
		BaseData: entities.BaseData{
			Title:      "Inhoud",
			ParentPath: "content",
		},
		Content: content,
	})
}

func (h *Handler) handleContentFormGET(c *gin.Context) {
	// Fetch content
	content, err := h.contentService.GetContent(c.Param(paramContentName))
	if err != nil {
		redirectWithMessage(c, sessions.Default(c), entities.MessageTypeError, err.Error(), "content/")
		return
	}

	// Render page
	c.HTML(http.StatusOK, "contentForm", entities.ContentFormData{
		BaseData: entities.BaseData{
			Title:      "Inhoud aanpassen",
			ParentPath: "content",
		},
		ContentName:   content.Name,
		IsHTMLContent: content.ContentType == baseEntities.ContentTypeHTML,
		Content: entities.Content{
			BodySimple: content.Body,
			BodyHTML:   template.HTML(content.Body),
		},
	})
}

func (h *Handler) handleContentFormPOST(c *gin.Context) {
	// Parse body
	contentName := c.Param(paramContentName)
	content := entities.Content{}
	err := c.MustBindWith(&content, binding.FormPost)
	if err != nil {
		renderContentFormWithError(c, content, contentName, true, fmt.Sprintf("Ongeldige data ontvangen: %v", err))
		return
	}

	// Fetch content
	currentContent, err := h.contentService.GetContent(contentName)
	if err != nil {
		renderContentFormWithError(c, content, contentName, true, fmt.Sprintf("Inhoud %s niet gevonden: %v", contentName, err))
		return
	}

	// Update content
	switch currentContent.ContentType {
	case baseEntities.ContentTypeSimple:
		currentContent.Body = content.BodySimple
	case baseEntities.ContentTypeHTML:
		currentContent.Body = string(content.BodyHTML)
	}
	_, err = h.contentService.UpdateContent(currentContent)
	if err != nil {
		isHTMLContent := currentContent.ContentType == baseEntities.ContentTypeHTML
		renderContentFormWithError(c, content, contentName, isHTMLContent, fmt.Sprintf("Aanpassen van inhoud mislukt: %v", err))
		return
	}

	// Update successful
	redirectWithMessage(c, sessions.Default(c), entities.MessageTypeSuccess, "Inhoud successvol aangepast.", "content/")
}

func renderContentFormWithError(c *gin.Context, content entities.Content, contentName string, isHTMLContent bool, message string) {
	c.HTML(http.StatusOK, "contentForm", &entities.ContentFormData{
		BaseData: entities.BaseData{
			Title:      "Inhoud aanpassen",
			ParentPath: "content",
			Messages: []entities.Message{{
				Type:    entities.MessageTypeError,
				Content: message,
			}},
		},
		Content:       content,
		ContentName:   contentName,
		IsHTMLContent: isHTMLContent,
	})
}
