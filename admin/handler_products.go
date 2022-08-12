package admin

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
	baseEntities "github.com/JenswBE/go-commerce/entities"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

const paramProductID = "product_id"

func (h *Handler) handleProductsList(c *gin.Context) {
	// Define base data
	baseData := entities.BaseData{
		Title:      "Producten",
		ParentPath: "products",
	}

	// Fetch products
	products, err := h.productService.ListProducts(nil)
	if err != nil {
		baseData.AddMessage(entities.MessageTypeError, "Ophalen van producten mislukt: %v", err)
		html(c, http.StatusInternalServerError, &entities.ProductsListTemplate{BaseData: baseData})
		return
	}

	// Fetch manufacturers
	manufacturers, err := h.productService.ListManufacturers(nil)
	if err != nil {
		baseData.AddMessage(entities.MessageTypeError, "Ophalen van merken voor producten mislukt: %v", err)
		html(c, http.StatusInternalServerError, &entities.ProductsListTemplate{BaseData: baseData})
		return
	}
	manufacturersMap := make(map[uuid.UUID]baseEntities.Manufacturer, len(manufacturers))
	for _, manufacturer := range manufacturers {
		manufacturersMap[manufacturer.ID] = *manufacturer
	}

	// Generate public URL's
	publicURLMap := make(map[uuid.UUID]string, len(products))
	if h.features.Products.PublicURLTemplate != "" {
		for _, product := range products {
			urlBuilder := strings.Builder{}
			err = h.features.Products.PublicURLTemplateParsed.Execute(&urlBuilder, product)
			if err != nil {
				baseData.AddMessage(entities.MessageTypeError, "Opbouwen van publieke link mislukt voor product %s: %v", product.Name, err)
				html(c, http.StatusInternalServerError, &entities.ProductsListTemplate{BaseData: baseData})
				return
			}
			publicURLMap[product.ID] = urlBuilder.String()
		}
	}

	// Render page
	htmlWithFlashes(c, http.StatusOK, &entities.ProductsListTemplate{
		BaseData:         baseData,
		Products:         products,
		ManufacturersMap: manufacturersMap,
		PublicURLMap:     publicURLMap,
	})
}

func (h *Handler) handleProductsFormGET(c *gin.Context) {
	// Check if new product
	paramID := c.Param(paramProductID)
	if paramID == "new" {
		html(c, http.StatusOK, &entities.ProductsFormTemplate{
			BaseData: entities.BaseData{
				Title:      productsFormTitle(true),
				ParentPath: "products",
			},
			IsNew: true,
		})
		return
	}

	// Get session
	session := sessions.Default(c)

	// Parse ID parameter
	id, err := parseUUID(paramID, i18n.ObjectTypeProduct)
	if err != nil {
		log.Debug().Err(err).Str("product_id", paramID).Msg("Invalid product ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "products/")
		return
	}

	// Fetch product
	product, err := h.productService.GetProduct(id, true, nil)
	if err != nil {
		log.Debug().Err(err).Str("product_id", paramID).Msg("Failed to fetch product")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "products/")
		return
	}

	// Render page
	html(c, http.StatusOK, &entities.ProductsFormTemplate{
		BaseData: entities.BaseData{
			Title:      productsFormTitle(false),
			ParentPath: "products",
		},
		IsNew:   false,
		Product: entities.ProductFromEntity(product),
	})
}

func (h *Handler) handleProductsFormPOST(c *gin.Context) {
	// Check if new product
	paramID := c.Param(paramProductID)
	isNew := paramID == "new"

	// Parse body
	product := entities.Product{}
	err := c.MustBindWith(&product, binding.FormPost)
	if err != nil {
		renderProductsFormWithError(c, isNew, product, fmt.Sprintf("Ongeldige data ontvangen: %v", err))
		return
	}

	// Create new entity
	productEntity := product.ToEntity()
	if isNew {
		_, err := h.productService.CreateProduct(&productEntity)
		if err != nil {
			renderProductsFormWithError(c, isNew, product, fmt.Sprintf("Toevoegen van product mislukt: %v", err))
			return
		}
	} else {
		// Parse ID parameter
		productEntity.ID, err = parseUUID(paramID, i18n.ObjectTypeProduct)
		if err != nil {
			renderProductsFormWithError(c, isNew, product, fmt.Sprintf("Ongeldige product ID %s: %v", paramID, err))
			return
		}

		// Update product
		_, err := h.productService.UpdateProduct(&productEntity)
		if err != nil {
			renderProductsFormWithError(c, isNew, product, fmt.Sprintf("Aanpassen van product mislukt: %v", err))
			return
		}
	}

	// Upsert successful
	redirectWithMessage(c, sessions.Default(c), entities.MessageTypeSuccess, "Product successvol toegevoegd/aangepast.", "products/")
}

func productsFormTitle(isNew bool) string {
	if isNew {
		return "Product toevoegen"
	}
	return "Product aanpassen"
}

func renderProductsFormWithError(c *gin.Context, isNew bool, product entities.Product, message string) {
	html(c, http.StatusOK, &entities.ProductsFormTemplate{
		BaseData: entities.BaseData{
			Title:      productsFormTitle(isNew),
			ParentPath: "products",
			Messages: []entities.Message{{
				Type:    entities.MessageTypeError,
				Content: message,
			}},
		},
		IsNew:   isNew,
		Product: product,
	})
}

func (h *Handler) handleProductsDelete(c *gin.Context) {
	// Get session
	session := sessions.Default(c)

	// Parse parameters
	id, err := parseUUID(c.Param(paramProductID), i18n.ObjectTypeProduct)
	if err != nil {
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "products/")
		return
	}

	// Call service
	err = h.productService.DeleteProduct(id)
	if err != nil {
		msg := i18n.DeleteFailed(i18n.ObjectTypeProduct, id, err)
		redirectWithMessage(c, session, entities.MessageTypeError, msg, "products/")
		return
	}

	// Call successful
	msg := i18n.DeleteSuccessful(i18n.ObjectTypeProduct)
	redirectWithMessage(c, session, entities.MessageTypeSuccess, msg, "products/")
}
