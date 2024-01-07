package admin

import (
	"cmp"
	"fmt"
	"net/http"
	"path"
	"slices"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
	baseEntities "github.com/JenswBE/go-commerce/entities"
)

const (
	paramServiceID         = "service_id"
	paramServiceCategoryID = "service_category_id"
)

func (h *Handler) handleServiceCategoriesList(c *gin.Context) {
	// Define base data
	baseData := entities.BaseData{
		Title:      "Dienstsoorten",
		ParentPath: "service_categories",
	}

	// Fetch service categories
	serviceCategories, err := h.productService.ListServiceCategories(false)
	if err != nil {
		baseData.AddMessage(entities.MessageTypeError, "Ophalen van dienstsoorten mislukt: %v", err)
		html(c, http.StatusOK, &entities.ServiceCategoriesListTemplate{BaseData: baseData})
		return
	}

	// Render page
	htmlWithFlashes(c, &entities.ServiceCategoriesListTemplate{
		BaseData:          baseData,
		ServiceCategories: unresolveServiceCategories(serviceCategories),
	})
}

func unresolveServiceCategories(resolved []*baseEntities.ResolvedServiceCategory) []*baseEntities.ServiceCategory {
	return lo.Map(resolved, func(sc *baseEntities.ResolvedServiceCategory, _ int) *baseEntities.ServiceCategory {
		return &sc.ServiceCategory
	})
}

func (h *Handler) handleServiceCategoriesUpdateOrder(c *gin.Context) {
	// Init handler
	session := sessions.Default(c)
	serviceCategoryParamID := c.Param(paramServiceCategoryID)
	handlerLog := log.With().Str("service_category_id", serviceCategoryParamID).Logger()

	// Parse service category ID parameter
	serviceCategoryID, err := parseID(serviceCategoryParamID, i18n.ObjectTypeServiceCategory)
	if err != nil {
		handlerLog.Debug().Err(err).Msg("Invalid service category ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "/service_categories")
		return
	}

	// Parse body
	newOrderString, ok := c.GetPostForm("new_order")
	if !ok {
		handlerLog.Debug().Err(err).Msg("Form field new_order missing")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "/service_categories")
		return
	}
	newOrder, err := strconv.Atoi(newOrderString)
	if err != nil {
		handlerLog.Debug().Err(err).Str("new_order", newOrderString).Msg("new_order is not a number")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "/service_categories")
		return
	}

	// Fetch service category
	serviceCategory, err := h.productService.GetServiceCategory(serviceCategoryID, false)
	if err != nil {
		handlerLog.Debug().Err(err).Msg("Failed to get service category to update order")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "/service_categories")
		return
	}

	// Update category order
	serviceCategory.Order = newOrder
	_, err = h.productService.UpdateServiceCategory(&serviceCategory.ServiceCategory)
	if err != nil {
		handlerLog.Debug().Err(err).Int("new_order", newOrder).Msg("Failed update order of service category")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "/service_categories")
		return
	}

	// Render page
	redirect(c, "/service_categories")
}

func (h *Handler) handleServiceCategoriesFormGET(c *gin.Context) {
	// Check if new category
	paramID := c.Param(paramServiceCategoryID)
	if paramID == IDNew {
		html(c, http.StatusOK, &entities.ServiceCategoriesFormTemplate{
			BaseData: entities.BaseData{
				Title:      serviceCategoriesFormTitle(true),
				ParentPath: "service_categories",
			},
			IsNew: true,
		})
		return
	}

	// Get session
	session := sessions.Default(c)

	// Parse ID parameter
	id, err := parseID(paramID, i18n.ObjectTypeServiceCategory)
	if err != nil {
		log.Debug().Err(err).Str("service_category_id", paramID).Msg("Invalid service category ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "service_categories/")
		return
	}

	// Fetch service category
	serviceCategory, err := h.productService.GetServiceCategory(id, false)
	if err != nil {
		log.Debug().Err(err).Str("service_category_id", paramID).Msg("Failed to fetch service category")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "service_categories/")
		return
	}

	// Render page
	html(c, http.StatusOK, &entities.ServiceCategoriesFormTemplate{
		BaseData: entities.BaseData{
			Title:      serviceCategoriesFormTitle(false),
			ParentPath: "service_categories",
		},
		IsNew:           false,
		ServiceCategory: entities.ServiceCategoryFromEntity(&serviceCategory.ServiceCategory),
	})
}

func (h *Handler) handleServiceCategoriesFormPOST(c *gin.Context) {
	// Check if new service category
	paramID := c.Param(paramServiceCategoryID)
	isNew := paramID == IDNew

	// Parse body
	serviceCategory := entities.ServiceCategory{}
	err := c.MustBindWith(&serviceCategory, binding.FormPost)
	if err != nil {
		renderServiceCategoriesFormWithError(c, isNew, serviceCategory, fmt.Sprintf("Ongeldige data ontvangen: %v", err))
		return
	}

	// Create new entity
	serviceCategoryEntity := serviceCategory.ToEntity()
	if isNew {
		_, err := h.productService.CreateServiceCategory(&serviceCategoryEntity)
		if err != nil {
			renderServiceCategoriesFormWithError(c, isNew, serviceCategory, fmt.Sprintf("Toevoegen van dienstsoort mislukt: %v", err))
			return
		}
	} else {
		// Parse ID parameter
		serviceCategoryID, err := parseID(paramID, i18n.ObjectTypeServiceCategory)
		if err != nil {
			renderServiceCategoriesFormWithError(c, isNew, serviceCategory, fmt.Sprintf("Ongeldige dienstsoort ID %s: %v", paramID, err))
			return
		}

		// Fetch category
		current, err := h.productService.GetServiceCategory(serviceCategoryID, false)
		if err != nil {
			renderServiceCategoriesFormWithError(c, isNew, serviceCategory, fmt.Sprintf("Dienstsoort %s niet gevonden: %v", paramID, err))
			return
		}

		// Update category
		current.Name = serviceCategoryEntity.Name
		_, err = h.productService.UpdateServiceCategory(&current.ServiceCategory)
		if err != nil {
			renderServiceCategoriesFormWithError(c, isNew, serviceCategory, fmt.Sprintf("Aanpassen van dienstsoort mislukt: %v", err))
			return
		}
	}

	// Upsert successful
	redirectWithMessage(c, sessions.Default(c), entities.MessageTypeSuccess, "Dienstsoort successvol toegevoegd/aangepast.", "service_categories/")
}

func serviceCategoriesFormTitle(isNew bool) string {
	if isNew {
		return "Dienstsoort toevoegen"
	}
	return "Dienstsoort aanpassen"
}

func renderServiceCategoriesFormWithError(c *gin.Context, isNew bool, serviceCategory entities.ServiceCategory, message string) {
	html(c, http.StatusOK, &entities.ServiceCategoriesFormTemplate{
		BaseData: entities.BaseData{
			Title:      serviceCategoriesFormTitle(isNew),
			ParentPath: "service_categories",
			Messages: []entities.Message{{
				Type:    entities.MessageTypeError,
				Content: message,
			}},
		},
		IsNew:           isNew,
		ServiceCategory: serviceCategory,
	})
}

func (h *Handler) handleServiceCategoriesDelete(c *gin.Context) {
	// Get session
	session := sessions.Default(c)

	// Parse parameters
	id, err := parseID(c.Param(paramServiceCategoryID), i18n.ObjectTypeServiceCategory)
	if err != nil {
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "service_categories/")
		return
	}

	// Call service
	err = h.productService.DeleteServiceCategory(id)
	if err != nil {
		msg := i18n.DeleteFailed(i18n.ObjectTypeServiceCategory, "", err)
		redirectWithMessage(c, session, entities.MessageTypeError, msg, "service_categories/")
		return
	}

	// Call successful
	msg := i18n.DeleteSuccessful(i18n.ObjectTypeServiceCategory)
	redirectWithMessage(c, session, entities.MessageTypeSuccess, msg, "service_categories/")
}

func (h *Handler) handleServicesList(c *gin.Context) {
	// Get session
	session := sessions.Default(c)

	// Parse ID parameter
	paramID := c.Param(paramServiceCategoryID)
	id, err := parseID(paramID, i18n.ObjectTypeServiceCategory)
	if err != nil {
		log.Debug().Err(err).Str("service_category_id", paramID).Msg("Invalid service category ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "service_categories/")
		return
	}

	// Fetch resolved service category
	svcCat, err := h.productService.GetServiceCategory(id, true)
	if err != nil {
		log.Debug().Err(err).Str("service_category_id", paramID).Msg("Failed to fetch service category")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "service_categories/")
		return
	}

	// Render page
	htmlWithFlashes(c, &entities.ServicesListTemplate{
		BaseData: entities.BaseData{
			Title:      fmt.Sprintf(`Diensten in "%s"`, svcCat.Name),
			ParentPath: "service_categories",
		},
		Services:        svcCat.Services,
		ServiceCategory: &svcCat.ServiceCategory,
	})
}

func (h *Handler) handleServicesUpdateOrder(c *gin.Context) {
	// Init and get session
	paramSvcCatID := c.Param(paramServiceCategoryID)
	paramSvcID := c.Param(paramServiceID)
	logger := log.With().Str("service_category_id", paramSvcCatID).Str("service_id", paramSvcID).Logger()
	redirectURL := path.Join("service_categories", paramSvcCatID, "services")
	session := sessions.Default(c)

	// Parse service ID parameter
	serviceID, err := parseID(paramSvcID, i18n.ObjectTypeService)
	if err != nil {
		logger.Debug().Err(err).Msg("Invalid service ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), redirectURL)
		return
	}

	// Parse body
	newOrderString, ok := c.GetPostForm("new_order")
	if !ok {
		logger.Debug().Err(err).Msg("Form field new_order missing")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), redirectURL)
		return
	}
	newOrder, err := strconv.Atoi(newOrderString)
	if err != nil {
		logger.Debug().Err(err).Str("new_order", newOrderString).Msg("new_order is not a number")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), redirectURL)
		return
	}

	// Fetch service
	service, err := h.productService.GetService(serviceID)
	if err != nil {
		logger.Debug().Err(err).Msg("Failed to get service to update order")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), redirectURL)
		return
	}

	// Update service order
	service.Order = newOrder
	_, err = h.productService.UpdateService(service)
	if err != nil {
		logger.Debug().Err(err).Int("new_order", newOrder).Msg("Failed update order of service")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), redirectURL)
		return
	}

	// Render page
	redirect(c, redirectURL)
}

func (h *Handler) handleServicesFormGET(c *gin.Context) {
	// Init and get session
	paramSvcCatID := c.Param(paramServiceCategoryID)
	paramSvcID := c.Param(paramServiceID)
	logger := log.With().Str("service_category_id", paramSvcCatID).Str("service_id", paramSvcID).Logger()
	redirectURL := path.Join("service_categories", paramSvcCatID, "services")
	session := sessions.Default(c)

	// Fetch service categories
	resolvedServiceCategories, err := h.productService.ListServiceCategories(false)
	if err != nil {
		logger.Debug().Err(err).Msg("Failed to fetch service categories for service form")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), redirectURL)
		return
	}
	serviceCategories := unresolveServiceCategories(resolvedServiceCategories)
	slices.SortFunc(serviceCategories, func(a, b *baseEntities.ServiceCategory) int { return cmp.Compare(a.Name, b.Name) })

	// Check if new service
	if paramSvcID == IDNew {
		html(c, http.StatusOK, &entities.ServicesFormTemplate{
			BaseData: entities.BaseData{
				Title:      servicesFormTitle(true),
				ParentPath: "service_categories",
			},
			IsNew:             true,
			Service:           entities.Service{ServiceCategoryID: paramSvcCatID},
			ServiceCategories: serviceCategories,
		})
		return
	}

	// Parse service ID parameter
	id, err := parseID(paramSvcID, i18n.ObjectTypeService)
	if err != nil {
		logger.Debug().Err(err).Msg("Invalid service ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), redirectURL)
		return
	}

	// Fetch service
	service, err := h.productService.GetService(id)
	if err != nil {
		logger.Debug().Err(err).Msg("Failed to fetch service")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), redirectURL)
		return
	}

	// Render page
	html(c, http.StatusOK, &entities.ServicesFormTemplate{
		BaseData: entities.BaseData{
			Title:      servicesFormTitle(false),
			ParentPath: "service_categories",
		},
		IsNew:             false,
		Service:           entities.ServiceFromEntity(service),
		ServiceCategories: serviceCategories,
	})
}

func (h *Handler) handleServicesFormPOST(c *gin.Context) {
	// Init and get session
	paramSvcCatID := c.Param(paramServiceCategoryID)
	paramSvcID := c.Param(paramServiceID)
	logger := log.With().Str("service_category_id", paramSvcCatID).Str("service_id", paramSvcID).Logger()
	redirectURL := path.Join("service_categories", paramSvcCatID, "services")
	session := sessions.Default(c)

	// Check if new product
	isNew := paramSvcID == IDNew

	// Fetch service categories
	resolvedServiceCategories, err := h.productService.ListServiceCategories(false)
	if err != nil {
		logger.Debug().Err(err).Msg("Failed to fetch service categories for service form")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), redirectURL)
		return
	}
	serviceCategories := unresolveServiceCategories(resolvedServiceCategories)
	slices.SortFunc(serviceCategories, func(a, b *baseEntities.ServiceCategory) int { return cmp.Compare(a.Name, b.Name) })

	// Parse body
	service := entities.Service{}
	err = c.MustBindWith(&service, binding.FormPost)
	if err != nil {
		renderServicesFormWithError(c, isNew, service, serviceCategories, fmt.Sprintf("Ongeldige data ontvangen: %v", err))
		return
	}

	// Convert to entity
	serviceEntity, err := service.ToEntity()
	if err != nil {
		renderServicesFormWithError(c, isNew, service, serviceCategories, fmt.Sprintf("Ongeldige data ontvangen: %v", err))
		return
	}

	// Process upsert
	if isNew {
		_, err := h.productService.CreateService(&serviceEntity)
		if err != nil {
			renderServicesFormWithError(c, isNew, service, serviceCategories, fmt.Sprintf("Toevoegen van dienst mislukt: %v", err))
			return
		}
	} else {
		// Parse ID parameter
		serviceEntity.ID, err = parseID(paramSvcID, i18n.ObjectTypeService)
		if err != nil {
			renderServicesFormWithError(c, isNew, service, serviceCategories, fmt.Sprintf("Ongeldige dienst ID %s: %v", paramSvcID, err))
			return
		}

		// Update service
		_, err := h.productService.UpdateService(&serviceEntity)
		if err != nil {
			renderServicesFormWithError(c, isNew, service, serviceCategories, fmt.Sprintf("Aanpassen van dienst mislukt: %v", err))
			return
		}
	}

	// Upsert successful
	redirectWithMessage(c, sessions.Default(c), entities.MessageTypeSuccess, "Dienst successvol toegevoegd/aangepast.", redirectURL)
}

func servicesFormTitle(isNew bool) string {
	if isNew {
		return "Dienst toevoegen"
	}
	return "Dienst aanpassen"
}

func renderServicesFormWithError(c *gin.Context, isNew bool, service entities.Service, serviceCategories []*baseEntities.ServiceCategory, message string) {
	html(c, http.StatusOK, &entities.ServicesFormTemplate{
		BaseData: entities.BaseData{
			Title:      servicesFormTitle(isNew),
			ParentPath: "service_categories",
			Messages: []entities.Message{{
				Type:    entities.MessageTypeError,
				Content: message,
			}},
		},
		IsNew:             isNew,
		Service:           service,
		ServiceCategories: serviceCategories,
	})
}

func (h *Handler) handleServicesDelete(c *gin.Context) {
	// Init and get session
	paramSvcCatID := c.Param(paramServiceCategoryID)
	paramSvcID := c.Param(paramServiceID)
	redirectURL := path.Join("service_categories", paramSvcCatID, "services")
	session := sessions.Default(c)

	// Parse parameters
	id, err := parseID(paramSvcID, i18n.ObjectTypeService)
	if err != nil {
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), redirectURL)
		return
	}

	// Call service
	err = h.productService.DeleteService(id)
	if err != nil {
		msg := i18n.DeleteFailed(i18n.ObjectTypeService, "", err)
		redirectWithMessage(c, session, entities.MessageTypeError, msg, redirectURL)
		return
	}

	// Call successful
	msg := i18n.DeleteSuccessful(i18n.ObjectTypeService)
	redirectWithMessage(c, session, entities.MessageTypeSuccess, msg, redirectURL)
}
