package admin

import (
	"net/http"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const paramManufacturerID = "manufacturer_id"

func (h *AdminHandler) handleManufacturersList(c *gin.Context) {
	htmlWithFlashes(c, http.StatusOK, "manufacturersList", &entities.ProductsListData{
		BaseData: entities.BaseData{
			Title:      "Merken",
			ParentPath: "manufacturers",
		},
	})
}

func (h *AdminHandler) handleManufacturersEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "manufacturersList", entities.ProductsListData{
		BaseData: entities.BaseData{
			Title:      "Merken",
			ParentPath: "manufacturers",
		},
	})
}

func (h *AdminHandler) handleManufacturersDelete(c *gin.Context) {
	// Get session
	session := sessions.Default(c)

	// Parse parameters
	manID, err := parseUUID(c.Param(paramManufacturerID), objectTypeManufacturer)
	if err != nil {
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "manufacturers/")
		return
	}

	// Call service
	err = h.productService.DeleteManufacturer(manID)
	if err != nil {
		msg := i18n.DeleteFailed(i18n.ObjectTypeManufacturer, manID, err)
		redirectWithMessage(c, session, entities.MessageTypeError, msg, "manufacturers/")
		return
	}

	// Call successful
	msg := i18n.DeleteSuccessful(i18n.ObjectTypeManufacturer)
	redirectWithMessage(c, session, entities.MessageTypeSuccess, msg, "manufacturers/")
}
