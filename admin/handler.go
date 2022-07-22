package admin

import (
	"errors"
	"net/http"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
	"github.com/JenswBE/go-commerce/usecases/product"
	"github.com/JenswBE/go-commerce/utils/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const PrefixAdmin = "/admin/"

type AdminHandler struct {
	productService product.Usecase
	middlewares    gin.HandlersChain
}

func NewAdminGUIHandler(productService product.Usecase, sessionAuthKey [64]byte, authVerifier auth.Verifier) (*AdminHandler, error) {
	return &AdminHandler{
		productService: productService,
		middlewares: gin.HandlersChain{
			sessions.Sessions("gocom", cookie.NewStore(sessionAuthKey[:])),
			auth.NewAuthMiddleware(authVerifier).EnforceRoles([]string{auth.RoleAdmin}),
		},
	}, nil
}

func (h *AdminHandler) RegisterRoutes(r *gin.Engine) {
	// Register middlewares
	rg := r.Group(PrefixAdmin)
	rg.Use(h.middlewares...)

	// Register static routes
	rg.Static("static", "admin/html/static")

	// Register dynamic routes
	rg.GET("/", func(c *gin.Context) { c.Redirect(http.StatusTemporaryRedirect, "products/") })
	// rg.GET("categories/", handleCategoriesList)
	rg.Any("login/", h.handleLogin)
	rg.GET("logout/", h.handleLogout)
	rg.GET("manufacturers/", h.handleManufacturersList)
	rg.GET("manufacturers/:manufacturer_id/", h.handleManufacturersEdit)
	rg.Any("manufacturers/:manufacturer_id/delete/", h.handleManufacturersDelete)
	// rg.GET("products/", handleProductsList)
	// rg.GET("products/:product_id/", handleProductsEdit)
}

const (
	objectTypeManufacturer = "merk"
)

func parseUUID(input string, objectType i18n.ObjectType) (uuid.UUID, error) {
	output, err := uuid.Parse(input)
	if err != nil {
		return uuid.Nil, errors.New(i18n.InvalidUUID(objectType, input))
	}
	return output, nil
}

func htmlWithFlashes(c *gin.Context, code int, name string, obj entities.WithBaseData) {
	// Get and convert flashes
	session := sessions.Default(c)
	flashes := session.Flashes()
	messages := make([]entities.Message, 0, len(flashes))
	for _, flash := range flashes {
		messages = append(messages, entities.ParseMessage(flash.(string)))
	}
	obj.SetMessages(messages)

	// Save session
	err := session.Save()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to save session after retrieving flashes: %v. Flashes: %v", err, flashes)
		return
	}

	// Display page
	c.HTML(code, name, obj)
}

func redirectWithMessage(c *gin.Context, session sessions.Session, messageType entities.MessageType, message, adminRedirectLocation string) {
	session.AddFlash(entities.Message{
		Type:    messageType,
		Content: message,
	}.String())
	err := session.Save()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to save flash in session: %v", err)
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, PrefixAdmin+adminRedirectLocation)
}
