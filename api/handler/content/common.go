package content

import (
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/usecases/content"
	"github.com/gin-gonic/gin"
)

const pathPrefixContent = "/content"
const pathPrefixEvents = "/events"

type ContentHandler struct {
	presenter *presenter.Presenter
	service   content.Usecase
}

func NewContentHandler(p *presenter.Presenter, service content.Usecase) *ContentHandler {
	return &ContentHandler{
		presenter: p,
		service:   service,
	}
}

func (h *ContentHandler) RegisterPublicRoutes(rg *gin.RouterGroup) {
	groupContent := rg.Group(pathPrefixContent)
	groupContent.GET("/:content_name/", h.getContent)

	groupEvents := rg.Group(pathPrefixEvents)
	groupEvents.GET("/", h.listEvents)
	groupEvents.GET("/:id/", h.getEvent)
}

func (h *ContentHandler) RegisterAdminRoutes(rg *gin.RouterGroup) {
	groupContent := rg.Group(pathPrefixContent)
	groupContent.GET("/", h.listContent)
	groupContent.PUT("/:content_name/", h.updateContent)

	groupEvents := rg.Group(pathPrefixEvents)
	groupEvents.POST("/", h.createEvent)
	groupEvents.PUT("/:id/", h.updateEvent)
	groupEvents.DELETE("/:id/", h.deleteEvent)
}
