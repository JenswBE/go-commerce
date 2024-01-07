package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/zerolog/log"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
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
	htmlWithFlashes(c, &entities.CategoriesListTemplate{
		BaseData:   baseData,
		Categories: categories,
	})
}

func (h *Handler) handleCategoriesUpdateOrder(c *gin.Context) {
	// Init handler
	session := sessions.Default(c)
	categoryParamID := c.Param(paramCategoryID)
	handlerLog := log.With().Str("category_id", categoryParamID).Logger()

	// Parse category ID parameter
	categoryID, err := parseID(categoryParamID, i18n.ObjectTypeCategory)
	if err != nil {
		handlerLog.Debug().Err(err).Msg("Invalid category ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "/categories")
		return
	}

	// Parse body
	newOrderString, ok := c.GetPostForm("new_order")
	if !ok {
		handlerLog.Debug().Err(err).Msg("Form field new_order missing")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "/categories")
		return
	}
	newOrder, err := strconv.Atoi(newOrderString)
	if err != nil {
		handlerLog.Debug().Err(err).Str("new_order", newOrderString).Msg("new_order is not a number")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "/categories")
		return
	}

	// Fetch category
	category, err := h.productService.GetCategory(categoryID, nil)
	if err != nil {
		handlerLog.Debug().Err(err).Msg("Failed to get category to update order")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "/categories")
		return
	}

	// Update category order
	category.Order = newOrder
	_, err = h.productService.UpdateCategory(category)
	if err != nil {
		handlerLog.Debug().Err(err).Int("new_order", newOrder).Msg("Failed update order of category")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "/categories")
		return
	}

	// Render page
	redirect(c, "/categories")
}

func (h *Handler) handleCategoriesFormGET(c *gin.Context) {
	// Check if new category
	paramID := c.Param(paramCategoryID)
	if paramID == IDNew {
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
	id, err := parseID(paramID, i18n.ObjectTypeCategory)
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
	isNew := paramID == IDNew

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
		categoryID, err := parseID(paramID, i18n.ObjectTypeCategory)
		if err != nil {
			renderCategoriesFormWithError(c, isNew, category, fmt.Sprintf("Ongeldige categorie ID %s: %v", paramID, err))
			return
		}

		// Fetch category
		current, err := h.productService.GetCategory(categoryID, nil)
		if err != nil {
			renderCategoriesFormWithError(c, isNew, category, fmt.Sprintf("Categorie %s niet gevonden: %v", paramID, err))
			return
		}

		// Update category
		current.Name = categoryEntity.Name
		current.Description = categoryEntity.Description
		_, err = h.productService.UpdateCategory(current)
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
	return "Categorie aanpassen"
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
	id, err := parseID(c.Param(paramCategoryID), i18n.ObjectTypeCategory)
	if err != nil {
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "categories/")
		return
	}

	// Call service
	err = h.productService.DeleteCategory(id)
	if err != nil {
		if strings.Contains(err.Error(), "foreign") {
			err = fmt.Errorf("deze categorie wordt nog gebruikt door producten. Verwijder eerst de categorie van de producten voordat je de categorie zelf verwijderd. (%w)", err)
		}
		msg := i18n.DeleteFailed(i18n.ObjectTypeCategory, "", err)
		redirectWithMessage(c, session, entities.MessageTypeError, msg, "categories/")
		return
	}

	// Call successful
	msg := i18n.DeleteSuccessful(i18n.ObjectTypeCategory)
	redirectWithMessage(c, session, entities.MessageTypeSuccess, msg, "categories/")
}
