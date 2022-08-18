package admin

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/zerolog/log"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
)

const paramManufacturerID = "manufacturer_id"

func (h *Handler) handleManufacturersList(c *gin.Context) {
	// Define base data
	baseData := entities.BaseData{
		Title:      "Merken",
		ParentPath: "manufacturers",
	}

	// Fetch manufacturers
	manufacturers, err := h.productService.ListManufacturers(ProductImageConfigs)
	if err != nil {
		baseData.AddMessage(entities.MessageTypeError, "Ophalen van merken mislukt: %v", err)
		html(c, http.StatusOK, &entities.ManufacturersListTemplate{BaseData: baseData})
		return
	}

	// Render page
	htmlWithFlashes(c, http.StatusOK, &entities.ManufacturersListTemplate{
		BaseData:      baseData,
		Manufacturers: manufacturers,
	})
}

func (h *Handler) handleManufacturersFormGET(c *gin.Context) {
	// Check if new manufacturer
	paramID := c.Param(paramManufacturerID)
	if paramID == "new" {
		html(c, http.StatusOK, &entities.ManufacturersFormTemplate{
			BaseData: entities.BaseData{
				Title:      manufacturersFormTitle(true),
				ParentPath: "manufacturers",
			},
			IsNew: true,
		})
		return
	}

	// Get session
	session := sessions.Default(c)

	// Parse ID parameter
	id, err := parseID(paramID, i18n.ObjectTypeManufacturer)
	if err != nil {
		log.Debug().Err(err).Str("manufacturer_id", paramID).Msg("Invalid manufacturer ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "manufacturers/")
		return
	}

	// Fetch manufacturer
	manufacturer, err := h.productService.GetManufacturer(id, nil)
	if err != nil {
		log.Debug().Err(err).Str("manufacturer_id", paramID).Msg("Failed to fetch manufacturer")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "manufacturers/")
		return
	}

	// Render page
	html(c, http.StatusOK, &entities.ManufacturersFormTemplate{
		BaseData: entities.BaseData{
			Title:      manufacturersFormTitle(false),
			ParentPath: "manufacturers",
		},
		IsNew:        false,
		Manufacturer: entities.ManufacturerFromEntity(manufacturer),
	})
}

func (h *Handler) handleManufacturersFormPOST(c *gin.Context) {
	// Check if new manufacturer
	paramID := c.Param(paramManufacturerID)
	isNew := paramID == "new"

	// Parse body
	manufacturer := entities.Manufacturer{}
	err := c.MustBindWith(&manufacturer, binding.FormPost)
	if err != nil {
		renderManufacturersFormWithError(c, isNew, manufacturer, fmt.Sprintf("Ongeldige data ontvangen: %v", err))
		return
	}

	// Create new entity
	manufacturerEntity := manufacturer.ToEntity()
	if isNew {
		_, err := h.productService.CreateManufacturer(&manufacturerEntity)
		if err != nil {
			renderManufacturersFormWithError(c, isNew, manufacturer, fmt.Sprintf("Toevoegen van merk mislukt: %v", err))
			return
		}
	} else {
		// Parse ID parameter
		manufacturerEntity.ID, err = parseID(paramID, i18n.ObjectTypeManufacturer)
		if err != nil {
			renderManufacturersFormWithError(c, isNew, manufacturer, fmt.Sprintf("Ongeldige merk ID %s: %v", paramID, err))
			return
		}

		// Update manufacturer
		_, err := h.productService.UpdateManufacturer(&manufacturerEntity)
		if err != nil {
			renderManufacturersFormWithError(c, isNew, manufacturer, fmt.Sprintf("Aanpassen van merk mislukt: %v", err))
			return
		}
	}

	// Upsert successful
	redirectWithMessage(c, sessions.Default(c), entities.MessageTypeSuccess, "Merk successvol toegevoegd/aangepast.", "manufacturers/")
}

func manufacturersFormTitle(isNew bool) string {
	if isNew {
		return "Merk toevoegen"
	}
	return "Merk aanpassen"
}

func renderManufacturersFormWithError(c *gin.Context, isNew bool, manufacturer entities.Manufacturer, message string) {
	html(c, http.StatusOK, &entities.ManufacturersFormTemplate{
		BaseData: entities.BaseData{
			Title:      manufacturersFormTitle(isNew),
			ParentPath: "manufacturers",
			Messages: []entities.Message{{
				Type:    entities.MessageTypeError,
				Content: message,
			}},
		},
		IsNew:        isNew,
		Manufacturer: manufacturer,
	})
}

func (h *Handler) handleManufacturersDelete(c *gin.Context) {
	// Get session
	session := sessions.Default(c)

	// Parse parameters
	id, err := parseID(c.Param(paramManufacturerID), i18n.ObjectTypeManufacturer)
	if err != nil {
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "manufacturers/")
		return
	}

	// Call service
	err = h.productService.DeleteManufacturer(id)
	if err != nil {
		if strings.Contains(err.Error(), "foreign") {
			err = fmt.Errorf("dit merk wordt nog gebruikt door producten. Verwijder eerst het merk van de producten voordat je het merk zelf verwijderd. (%w)", err)
		}
		msg := i18n.DeleteFailed(i18n.ObjectTypeManufacturer, "", err)
		redirectWithMessage(c, session, entities.MessageTypeError, msg, "manufacturers/")
		return
	}

	// Call successful
	msg := i18n.DeleteSuccessful(i18n.ObjectTypeManufacturer)
	redirectWithMessage(c, session, entities.MessageTypeSuccess, msg, "manufacturers/")
}

func (h *Handler) handleManufacturersImageGET(c *gin.Context) {
	// Init handler
	session := sessions.Default(c)
	paramID := c.Param(paramManufacturerID)
	handlerLog := log.With().Str("manufacturer_id", paramID).Logger()

	// Parse ID parameter
	id, err := parseID(paramID, i18n.ObjectTypeManufacturer)
	if err != nil {
		handlerLog.Debug().Err(err).Msg("Invalid manufacturer ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "manufacturers/")
		return
	}

	// Fetch product
	manufacturer, err := h.productService.GetManufacturer(id, ProductImageConfigs)
	if err != nil {
		handlerLog.Debug().Err(err).Msg("Failed to fetch manufacturer")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "manufacturers/")
		return
	}

	// Render page
	htmlWithFlashes(c, http.StatusOK, &entities.ManufacturersImageTemplate{
		BaseData: entities.BaseData{
			Title:      "Foto aanpassen",
			ParentPath: "manufacturers",
		},
		Manufacturer: *manufacturer,
	})
}

func (h *Handler) handleManufacturersImagePOST(c *gin.Context) {
	// Init handler
	session := sessions.Default(c)
	paramID := c.Param(paramManufacturerID)
	handlerLog := log.With().Str("manufacturer_id", paramID).Logger()

	// Parse ID parameter
	id, err := parseID(paramID, i18n.ObjectTypeManufacturer)
	if err != nil {
		handlerLog.Debug().Err(err).Msg("Invalid manufacturer ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), c.Request.URL.String())
		return
	}

	// Parse body
	images, err := parseFilesFromMultipart(c.Request)
	if err != nil {
		handlerLog.Debug().Err(err).Msg("Failed parse image files for manufacturer from multipart body")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), c.Request.URL.String())
		return
	}
	if len(images) != 1 {
		handlerLog.Debug().Err(err).Int("image_count", len(images)).Msg("Expected exactly 1 image file for manufacturer from multipart body")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), c.Request.URL.String())
		return
	}

	// Add images to product
	for imageName, imageBytes := range images {
		_, err = h.productService.UpsertManufacturerImage(id, imageName, imageBytes, ProductImageConfigs)
		if err != nil {
			handlerLog.Debug().Err(err).Str("product_id", paramID).Msg("Failed upsert image to manufacturer")
			redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), c.Request.URL.String())
			return
		}
	}

	// Render page
	c.Redirect(http.StatusSeeOther, c.Request.URL.String())
}

func (h *Handler) handleManufacturersImageDelete(c *gin.Context) {
	// Init handler
	session := sessions.Default(c)
	paramID := c.Param(paramManufacturerID)
	handlerLog := log.With().Str("manufacturer_id", paramID).Logger()

	// Parse manufacturer ID parameter
	id, err := parseID(paramID, i18n.ObjectTypeManufacturer)
	if err != nil {
		handlerLog.Debug().Err(err).Msg("Invalid manufacturer ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "/manufacturers")
		return
	}

	// Remove image to product
	err = h.productService.DeleteManufacturerImage(id)
	if err != nil {
		handlerLog.Debug().Err(err).Msg("Failed to delete image of manufacturer")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), fmt.Sprintf("/manufacturers/%s/image/", id))
		return
	}

	// Render page
	redirect(c, fmt.Sprintf("/manufacturers/%s/image/", id))
}
