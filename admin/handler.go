package admin

import (
	"net/http"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct{}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}

func (h *AdminHandler) RegisterRoutes(rg *gin.RouterGroup) {
	// Register static routes
	rg.Static("static", "admin/html/static")

	// Register dynamic routes
	rg.GET("/", func(c *gin.Context) { c.Redirect(http.StatusTemporaryRedirect, "products/") })
	rg.GET("categories/", handleCategoriesList)
	rg.GET("login/", handleLogin)
	rg.GET("products/", handleProductsList)
	rg.GET("products/:product_id", handleProductsEdit)
}

func handleCategoriesList(c *gin.Context) {
	c.HTML(200, "categoriesList", entities.ProductsListData{
		BaseData: entities.BaseData{
			Title:      "Categorieën",
			ParentPath: "categories",
		},
	})
}

func handleLogin(c *gin.Context) {
	c.HTML(200, "login", gin.H{
		"title": "Inloggen",
	})
}

func handleProductsList(c *gin.Context) {
	c.HTML(200, "productsList", entities.ProductsListData{
		BaseData: entities.BaseData{
			Title:      "Producten",
			ParentPath: "products",
		},
	})
}

func handleProductsEdit(c *gin.Context) {
	c.HTML(200, "productsList", gin.H{
		"title": "Html5 Template Engine",
	})
}
