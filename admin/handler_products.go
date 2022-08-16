package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
	baseEntities "github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/zerolog/log"
)

const (
	paramProductID = "product_id"
	paramImageID   = "image_id"
)

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
	manufacturersMap := make(map[baseEntities.ID]baseEntities.Manufacturer, len(manufacturers))
	for _, manufacturer := range manufacturers {
		manufacturersMap[manufacturer.ID] = *manufacturer
	}

	// Generate public URL's
	publicURLMap := make(map[baseEntities.ID]string, len(products))
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
	// Get session
	session := sessions.Default(c)

	// Fetch categories
	paramID := c.Param(paramProductID)
	categories, err := h.productService.ListCategories(nil)
	if err != nil {
		log.Debug().Err(err).Str("product_id", paramID).Msg("Failed to fetch categories for product form")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "products/")
		return
	}

	// Fetch manufacturers
	manufacturers, err := h.productService.ListManufacturers(nil)
	if err != nil {
		log.Debug().Err(err).Str("product_id", paramID).Msg("Failed to fetch manufacturers for product form")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "products/")
		return
	}

	// Check if new product
	if paramID == "new" {
		html(c, http.StatusOK, &entities.ProductsFormTemplate{
			BaseData: entities.BaseData{
				Title:      productsFormTitle(true),
				ParentPath: "products",
			},
			IsNew:         true,
			Product:       entities.Product{},
			Categories:    categories,
			Manufacturers: manufacturers,
		})
		return
	}

	// Parse ID parameter
	id, err := parseID(paramID, i18n.ObjectTypeProduct)
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
		IsNew:         false,
		Product:       entities.ProductFromEntity(product),
		Categories:    categories,
		Manufacturers: manufacturers,
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

	// Convert to entity
	productEntity, err := product.ToEntity()
	if err != nil {
		renderProductsFormWithError(c, isNew, product, fmt.Sprintf("Ongeldige data ontvangen: %v", err))
		return
	}

	// Process upsert
	if isNew {
		_, err := h.productService.CreateProduct(&productEntity)
		if err != nil {
			renderProductsFormWithError(c, isNew, product, fmt.Sprintf("Toevoegen van product mislukt: %v", err))
			return
		}
	} else {
		// Parse ID parameter
		productEntity.ID, err = parseID(paramID, i18n.ObjectTypeProduct)
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
	id, err := parseID(c.Param(paramProductID), i18n.ObjectTypeProduct)
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

var ProductImageConfigs map[string]imageproxy.ImageConfig = map[string]imageproxy.ImageConfig{"200": {
	Width:        200,
	Height:       200,
	ResizingType: imageproxy.ResizingTypeFit,
}}

func (h *Handler) handleProductsImagesGET(c *gin.Context) {
	// Get session
	session := sessions.Default(c)

	// Parse ID parameter
	paramID := c.Param(paramProductID)
	id, err := parseID(paramID, i18n.ObjectTypeProduct)
	if err != nil {
		log.Debug().Err(err).Str("product_id", paramID).Msg("Invalid product ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "products/")
		return
	}

	// Fetch product
	product, err := h.productService.GetProduct(id, true, ProductImageConfigs)
	if err != nil {
		log.Debug().Err(err).Str("product_id", paramID).Msg("Failed to fetch product")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "products/")
		return
	}

	// Render page
	htmlWithFlashes(c, http.StatusOK, &entities.ProductsImagesTemplate{
		BaseData: entities.BaseData{
			Title:      "Foto's aanpassen",
			ParentPath: "products",
		},
		Product: product.Product,
	})
}

func (h *Handler) handleProductsImagesPOST(c *gin.Context) {
	// Get session
	session := sessions.Default(c)

	// Parse ID parameter
	paramID := c.Param(paramProductID)
	id, err := parseID(paramID, i18n.ObjectTypeProduct)
	if err != nil {
		log.Debug().Err(err).Str("product_id", paramID).Msg("Invalid product ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), c.Request.URL.String())
		return
	}

	// Parse body
	images, err := parseFilesFromMultipart(c.Request)
	if err != nil {
		log.Debug().Err(err).Str("product_id", paramID).Msg("Failed parse image files for product from multipart body")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), c.Request.URL.String())
		return
	}

	// Add images to product
	_, err = h.productService.AddProductImages(id, images, ProductImageConfigs)
	if err != nil {
		log.Debug().Err(err).Str("product_id", paramID).Msg("Failed add images to product")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), c.Request.URL.String())
		return
	}

	// Render page
	c.Redirect(http.StatusSeeOther, c.Request.URL.String())
}

func (h *Handler) handleProductsImagesUpdateOrder(c *gin.Context) {
	// Init handler
	session := sessions.Default(c)
	productParamID := c.Param(paramProductID)
	imageParamID := c.Param(paramImageID)
	handlerLog := log.With().Str("product_id", productParamID).Str("image_id", imageParamID).Logger()

	// Parse product ID parameter
	productID, err := parseID(productParamID, i18n.ObjectTypeProduct)
	if err != nil {
		handlerLog.Debug().Err(err).Msg("Invalid product ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "/products")
		return
	}

	// Parse image ID parameter
	imageID, err := parseID(imageParamID, i18n.ObjectTypeProduct)
	if err != nil {
		handlerLog.Debug().Err(err).Msg("Invalid image ID provided for product")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), fmt.Sprintf("/products/%s/images/", productID))
		return
	}

	// Parse body
	newOrderString, ok := c.GetPostForm("new_order")
	if !ok {
		handlerLog.Debug().Err(err).Msg("handleProductsImagesUpdateOrder: Form field new_order missing")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), fmt.Sprintf("/products/%s/images/", productID))
		return
	}
	newOrder, err := strconv.Atoi(newOrderString)
	if err != nil {
		handlerLog.Debug().Err(err).Str("new_order", newOrderString).Msg("handleProductsImagesUpdateOrder: new_order is not a number")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), fmt.Sprintf("/products/%s/images/", productID))
		return
	}

	// Update product image
	_, err = h.productService.UpdateProductImage(productID, imageID, newOrder)
	if err != nil {
		handlerLog.Debug().Err(err).Int("new_order", newOrder).Msg("handleProductsImagesUpdateOrder: Failed update order of product image")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), fmt.Sprintf("/products/%s/images/", productID))
		return
	}

	// Render page
	redirect(c, fmt.Sprintf("/products/%s/images/", productID))
}

func (h *Handler) handleProductsImagesDelete(c *gin.Context) {
	// Init handler
	session := sessions.Default(c)
	productParamID := c.Param(paramProductID)
	imageParamID := c.Param(paramImageID)
	handlerLog := log.With().Str("product_id", productParamID).Str("image_id", imageParamID).Logger()

	// Parse product ID parameter
	productID, err := parseID(productParamID, i18n.ObjectTypeProduct)
	if err != nil {
		handlerLog.Debug().Err(err).Msg("Invalid product ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "/products")
		return
	}

	// Parse image ID parameter
	imageID, err := parseID(imageParamID, i18n.ObjectTypeProduct)
	if err != nil {
		handlerLog.Debug().Err(err).Msg("Invalid image ID provided for product")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), fmt.Sprintf("/products/%s/images/", productID))
		return
	}

	// Remove image to product
	err = h.productService.DeleteProductImage(productID, imageID)
	if err != nil {
		handlerLog.Debug().Err(err).Msg("Failed to delete image of product")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), fmt.Sprintf("/products/%s/images/", productID))
		return
	}

	// Render page
	redirect(c, fmt.Sprintf("/products/%s/images/", productID))
}
