package product

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/usecases/product"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
)

const (
	pathPrefixCategories        = "/categories"
	pathPrefixManufacturers     = "/manufacturers"
	pathPrefixProducts          = "/products"
	pathPrefixServiceCategories = "/service_categories"
)

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

func (h *ProductHandler) RegisterRoutes(rg *gin.RouterGroup) {
	groupCategories := rg.Group(pathPrefixCategories)
	groupCategories.GET("/", h.listCategories)
	groupCategories.GET("/:id/", h.getCategory)

	groupManufacturers := rg.Group(pathPrefixManufacturers)
	groupManufacturers.GET("/", h.listManufacturers)
	groupManufacturers.GET("/:id/", h.getManufacturer)

	groupProducts := rg.Group(pathPrefixProducts)
	groupProducts.GET("/", h.listProducts)
	groupProducts.GET("/:id/", h.getProduct)

	groupServiceCategories := rg.Group(pathPrefixServiceCategories)
	groupServiceCategories.GET("/", h.listServiceCategories)
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
