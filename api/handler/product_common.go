package handler

import (
	"errors"
	"io"
	"net/http"

	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/usecases/product"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/gin-gonic/gin"
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

func (h *ProductHandler) RegisterReadRoutes(rg *gin.RouterGroup) {
	groupCategories := rg.Group("/categories")
	groupCategories.GET("", h.listCategories)
	groupCategories.GET("/:id", h.getCategory)

	groupManufacturers := rg.Group("/manufacturers")
	groupManufacturers.GET("", h.listManufacturers)
	groupManufacturers.GET("/:id", h.getManufacturer)

	groupProducts := rg.Group("/products")
	groupProducts.GET("", h.listProducts)
	groupProducts.GET("/:id", h.getProduct)
	groupProducts.GET("/:id/images", h.listProductImages)
}

func (h *ProductHandler) RegisterWriteRoutes(rg *gin.RouterGroup) {
	groupCategories := rg.Group("/categories")
	groupCategories.POST("", h.createCategory)
	groupCategories.PUT("/:id", h.updateCategory)
	groupCategories.DELETE("/:id", h.deleteCategory)
	groupCategories.PUT("/:id/image", h.upsertCategoryImage)
	groupCategories.DELETE("/:id/image", h.deleteCategoryImage)

	groupManufacturers := rg.Group("/manufacturers")
	groupManufacturers.POST("", h.createManufacturer)
	groupManufacturers.PUT("/:id", h.updateManufacturer)
	groupManufacturers.DELETE("/:id", h.deleteManufacturer)
	groupManufacturers.PUT("/:id/image", h.upsertManufacturerImage)
	groupManufacturers.DELETE("/:id/image", h.deleteManufacturerImage)

	groupProducts := rg.Group("/products")
	groupProducts.POST("", h.createProduct)
	groupProducts.PUT("/:id", h.updateProduct)
	groupProducts.DELETE("/:id", h.deleteProduct)
	groupProducts.POST("/:id/images", h.addProductImages)
	groupProducts.PUT("/:id/images/:image_id", h.updateProductImage)
	groupProducts.DELETE("/:id/images/:image_id", h.deleteProductImage)
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

func parseImageConfigParams(c *gin.Context) (*imageproxy.ImageConfig, error) {
	// Extract params
	width := c.Query("img_w")
	height := c.Query("img_h")
	resizingType := c.Query("img_r")

	// Only process when any param is set
	if width == "" && height == "" && resizingType == "" {
		return nil, nil
	}

	// Set defaults
	if resizingType == "" {
		resizingType = string(imageproxy.ResizingTypeFit)
	}

	// Parse config
	return imageproxy.ParseImageConfig(width, height, resizingType)
}
