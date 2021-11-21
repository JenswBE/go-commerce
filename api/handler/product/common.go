package product

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/usecases/product"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/gin-gonic/gin"
)

const pathPrefixCategories = "/categories"
const pathPrefixManufacturers = "/manufacturers"
const pathPrefixProducts = "/products"

type ProductHandler struct {
	presenter *presenter.Presenter
	service   product.Usecase
}

func NewProductHandler(p *presenter.Presenter, service product.Usecase) *ProductHandler {
	return &ProductHandler{
		presenter: p,
		service:   service,
	}
}

func (h *ProductHandler) RegisterPublicRoutes(rg *gin.RouterGroup) {
	groupCategories := rg.Group(pathPrefixCategories)
	groupCategories.GET("/", h.listCategories)
	groupCategories.GET("/:id/", h.getCategory)

	groupManufacturers := rg.Group(pathPrefixManufacturers)
	groupManufacturers.GET("/", h.listManufacturers)
	groupManufacturers.GET("/:id/", h.getManufacturer)

	groupProducts := rg.Group(pathPrefixProducts)
	groupProducts.GET("/", h.listProducts)
	groupProducts.GET("/:id/", h.getProduct)
	groupProducts.GET("/:id/images/", h.listProductImages)
}

func (h *ProductHandler) RegisterAdminRoutes(rg *gin.RouterGroup) {
	groupCategories := rg.Group(pathPrefixCategories)
	groupCategories.POST("/", h.createCategory)
	groupCategories.PUT("/:id/", h.updateCategory)
	groupCategories.DELETE("/:id/", h.deleteCategory)
	groupCategories.PUT("/:id/image/", h.upsertCategoryImage)
	groupCategories.DELETE("/:id/image/", h.deleteCategoryImage)

	groupManufacturers := rg.Group(pathPrefixManufacturers)
	groupManufacturers.POST("/", h.createManufacturer)
	groupManufacturers.PUT("/:id/", h.updateManufacturer)
	groupManufacturers.DELETE("/:id/", h.deleteManufacturer)
	groupManufacturers.PUT("/:id/image/", h.upsertManufacturerImage)
	groupManufacturers.DELETE("/:id/image/", h.deleteManufacturerImage)

	groupProducts := rg.Group(pathPrefixProducts)
	groupProducts.POST("/", h.createProduct)
	groupProducts.PUT("/:id/", h.updateProduct)
	groupProducts.DELETE("/:id/", h.deleteProduct)
	groupProducts.POST("/:id/images/", h.addProductImages)
	groupProducts.PUT("/:id/images/:image_id/", h.updateProductImage)
	groupProducts.DELETE("/:id/images/:image_id/", h.deleteProductImage)
}

func parseFilesFromMultipart(req *http.Request) (map[string][]byte, error) {
	// Create reader from request
	reader, err := req.MultipartReader()
	if err != nil {
		return nil, err
	}

	// Parse images
	images := map[string][]byte{}
	for {
		// Parse part
		part, err := reader.NextPart()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return nil, err
			}
		}

		// Add to images
		imageBytes, err := io.ReadAll(part)
		if err != nil {
			return nil, err
		}
		images[part.FileName()] = imageBytes
	}

	// Parsing successful
	return images, nil
}

func parseImageConfigsParam(c *gin.Context) (map[string]imageproxy.ImageConfig, error) {
	// No params on empty string
	param := c.Query("img")
	if param == "" {
		return nil, nil
	}

	// Split config strings
	configStrings := strings.Split(param, ",")
	configs := make(map[string]imageproxy.ImageConfig, len(configStrings))
	for _, configString := range configStrings {
		// Split config into parts
		configParts := strings.Split(configString, "_")

		// Set defaults
		width := configParts[0]
		height := width
		resizingType := string(imageproxy.ResizingTypeFit)

		// Set height if defined
		if len(configParts) > 1 {
			height = configParts[1]
		}

		// Set resize type if defined
		if len(configParts) > 2 {
			resizingType = configParts[2]
		}

		// Parse into config
		var err error
		configs[configString], err = imageproxy.ParseImageConfig(width, height, resizingType)
		if err != nil {
			return nil, err
		}
	}

	// Parse successful
	return configs, nil
}
