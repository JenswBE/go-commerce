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

func AddProductReadRoutes(rg *gin.RouterGroup, p *presenter.Presenter, service product.Usecase) {
	groupCategories := rg.Group("/categories")
	groupCategories.GET("", listCategories(p, service))
	groupCategories.GET("/:id", getCategory(p, service))

	groupManufacturers := rg.Group("/manufacturers")
	groupManufacturers.GET("", listManufacturers(p, service))
	groupManufacturers.GET("/:id", getManufacturer(p, service))

	groupProducts := rg.Group("/products")
	groupProducts.GET("", listProducts(p, service))
	groupProducts.GET("/:id", getProduct(p, service))
	groupProducts.GET("/:id/images", listProductImages(p, service))
}

func AddProductWriteRoutes(rg *gin.RouterGroup, p *presenter.Presenter, service product.Usecase) {
	groupCategories := rg.Group("/categories")
	groupCategories.POST("", createCategory(p, service))
	groupCategories.PUT("/:id", updateCategory(p, service))
	groupCategories.DELETE("/:id", deleteCategory(p, service))
	groupCategories.PUT("/:id/image", upsertCategoryImage(p, service))
	groupCategories.DELETE("/:id/image", deleteCategoryImage(p, service))

	groupManufacturers := rg.Group("/manufacturers")
	groupManufacturers.POST("", createManufacturer(p, service))
	groupManufacturers.PUT("/:id", updateManufacturer(p, service))
	groupManufacturers.DELETE("/:id", deleteManufacturer(p, service))
	groupManufacturers.PUT("/:id/image", upsertManufacturerImage(p, service))
	groupManufacturers.DELETE("/:id/image", deleteManufacturerImage(p, service))

	groupProducts := rg.Group("/products")
	groupProducts.POST("", createProduct(p, service))
	groupProducts.PUT("/:id", updateProduct(p, service))
	groupProducts.DELETE("/:id", deleteProduct(p, service))
	groupProducts.POST("/:id/images", addProductImages(p, service))
	groupProducts.PUT("/:id/images/:image_id", updateProductImage(p, service))
	groupProducts.DELETE("/:id/images/:image_id", deleteProductImage(p, service))
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
