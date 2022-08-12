package admin

import (
	"fmt"
	"net/http"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/zerolog/log"
)

const paramCategoryID = "category_id"

func (h *Handler) handleCategoriesList(c *gin.Context) {
	// Define base data
	baseData := entities.BaseData{
		Title:      "Categorieën",
		ParentPath: "categories",
	}

	// Fetch categories
	categories, err := h.productService.ListCategories(nil)
	if err != nil {
		baseData.AddMessage(entities.MessageTypeError, "Ophalen van categorieën mislukt: %v", err)
		html(c, http.StatusOK, &entities.CategoriesListTemplate{BaseData: baseData})
		return
	}

	// Render page
	htmlWithFlashes(c, http.StatusOK, &entities.CategoriesListTemplate{
		BaseData:   baseData,
		Categories: categories,
	})
}

func (h *Handler) handleCategoriesFormGET(c *gin.Context) {
	// Check if new category
	paramID := c.Param(paramCategoryID)
	if paramID == "new" {
		html(c, http.StatusOK, &entities.CategoriesFormTemplate{
			BaseData: entities.BaseData{
				Title:      categoriesFormTitle(true),
				ParentPath: "categories",
			},
			IsNew: true,
		})
		return
	}

	// Get session
	session := sessions.Default(c)

	// Parse ID parameter
	id, err := parseUUID(paramID, i18n.ObjectTypeCategory)
	if err != nil {
		log.Debug().Err(err).Str("category_id", paramID).Msg("Invalid category ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "categories/")
		return
	}

	// Fetch category
	category, err := h.productService.GetCategory(id, nil)
	if err != nil {
		log.Debug().Err(err).Str("category_id", paramID).Msg("Failed to fetch category")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "categories/")
		return
	}

	// Render page
	html(c, http.StatusOK, &entities.CategoriesFormTemplate{
		BaseData: entities.BaseData{
			Title:      categoriesFormTitle(false),
			ParentPath: "categories",
		},
		IsNew:    false,
		Category: entities.CategoryFromEntity(category),
	})
}

func (h *Handler) handleCategoriesFormPOST(c *gin.Context) {
	// Check if new category
	paramID := c.Param(paramCategoryID)
	isNew := paramID == "new"

	// Parse body
	category := entities.Category{}
	err := c.MustBindWith(&category, binding.FormPost)
	if err != nil {
		renderCategoriesFormWithError(c, isNew, category, fmt.Sprintf("Ongeldige data ontvangen: %v", err))
		return
	}

	// Create new entity
	categoryEntity := category.ToEntity()
	if isNew {
		_, err := h.productService.CreateCategory(&categoryEntity)
		if err != nil {
			renderCategoriesFormWithError(c, isNew, category, fmt.Sprintf("Toevoegen van categorie mislukt: %v", err))
			return
		}
	} else {
		// Parse ID parameter
		categoryEntity.ID, err = parseUUID(paramID, i18n.ObjectTypeCategory)
		if err != nil {
			renderCategoriesFormWithError(c, isNew, category, fmt.Sprintf("Ongeldige categorie ID %s: %v", paramID, err))
			return
		}

		// Update category
		_, err := h.productService.UpdateCategory(&categoryEntity)
		if err != nil {
			renderCategoriesFormWithError(c, isNew, category, fmt.Sprintf("Aanpassen van categorie mislukt: %v", err))
			return
		}
	}

	// Upsert successful
	redirectWithMessage(c, sessions.Default(c), entities.MessageTypeSuccess, "Categorie successvol toegevoegd/aangepast.", "categories/")
}

func categoriesFormTitle(isNew bool) string {
	if isNew {
		return "Categorie toevoegen"
	}
	return "Categoryement aanpassen"
}

func renderCategoriesFormWithError(c *gin.Context, isNew bool, category entities.Category, message string) {
	html(c, http.StatusOK, &entities.CategoriesFormTemplate{
		BaseData: entities.BaseData{
			Title:      categoriesFormTitle(isNew),
			ParentPath: "categories",
			Messages: []entities.Message{{
				Type:    entities.MessageTypeError,
				Content: message,
			}},
		},
		IsNew:    isNew,
		Category: category,
	})
}

func (h *Handler) handleCategoriesDelete(c *gin.Context) {
	// Get session
	session := sessions.Default(c)

	// Parse parameters
	id, err := parseUUID(c.Param(paramCategoryID), i18n.ObjectTypeCategory)
	if err != nil {
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "categories/")
		return
	}

	// Call service
	err = h.productService.DeleteCategory(id)
	if err != nil {
		msg := i18n.DeleteFailed(i18n.ObjectTypeCategory, id, err)
		redirectWithMessage(c, session, entities.MessageTypeError, msg, "categories/")
		return
	}

	// Call successful
	msg := i18n.DeleteSuccessful(i18n.ObjectTypeCategory)
	redirectWithMessage(c, session, entities.MessageTypeSuccess, msg, "categories/")
}