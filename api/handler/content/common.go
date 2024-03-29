package content

import (
	"github.com/gin-gonic/gin"

	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/usecases/content"
)

const (
	pathPrefixContent = "/content"
	pathPrefixEvents  = "/events"
)

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

func (h *ContentHandler) RegisterRoutes(rg *gin.RouterGroup) {
	groupContent := rg.Group(pathPrefixContent)
	groupContent.GET("/:content_name/", h.getContent)

	groupEvents := rg.Group(pathPrefixEvents)
	groupEvents.GET("/", h.listEvents)
	groupEvents.GET("/:id/", h.getEvent)
}
