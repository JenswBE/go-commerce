package admin

import (
	"errors"
	"net/http"
	"time"

	"github.com/JenswBE/go-commerce/admin/auth"
	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
	"github.com/JenswBE/go-commerce/usecases/content"
	"github.com/JenswBE/go-commerce/usecases/product"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const PrefixAdmin = "/admin/"

const (
	objectTypeEvent        = "evenement"
	objectTypeManufacturer = "merk"
	pathLogin              = "login/"
)

type AdminHandler struct {
	contentService       content.Service
	productService       product.Usecase
	authVerifier         auth.Verifier
	sessionAuthenticator *auth.SessionAuthenticator
	sessionAuthKey       [64]byte
	sessionEncKey        [32]byte
}

func NewAdminGUIHandler(productService product.Usecase, authVerifier auth.Verifier, sessionAuthKey [64]byte, sessionEncKey [32]byte) *AdminHandler {
	handler := &AdminHandler{
		productService: productService,
		authVerifier:   authVerifier,
		sessionAuthKey: sessionAuthKey,
		sessionEncKey:  sessionEncKey,
	}
	if authVerifier != nil {
		handler.sessionAuthenticator = auth.NewSessionAuthenticator(time.Hour * 24 * 7)
	}
	return handler
}

func (h *AdminHandler) RegisterRoutes(r *gin.Engine) {
	// Register middlewares
	notAuthenticatedGroup := r.Group(PrefixAdmin)
	cookieStore := cookie.NewStore(h.sessionAuthKey[:], h.sessionEncKey[:])
	cookieStore.Options(sessions.Options{
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	notAuthenticatedGroup.Use(sessions.Sessions("gocom", cookieStore))
	rg := notAuthenticatedGroup.Group("")
	if h.sessionAuthenticator != nil {
		rg.Use(h.sessionAuthenticator.MW(PrefixAdmin + pathLogin))
	}

	// Register static routes
	rg.Static("static", "admin/html/static")

	// Register dynamic routes
	rg.GET("/", func(c *gin.Context) { c.Redirect(http.StatusTemporaryRedirect, "products/") })
	// rg.GET("categories/", handleCategoriesList)
	notAuthenticatedGroup.GET(pathLogin, h.handleLogin)
	notAuthenticatedGroup.POST(pathLogin, h.handleLogin)
	rg.GET("logout/", h.handleLogout)
	rg.GET("events/", h.handleEventsList)
	rg.GET("events/:event_id/", h.handleEventsForm)
	rg.POST("events/:event_id/", h.handleEventsForm)
	rg.POST("events/:event_id/delete/", h.handleEventsDelete)
	rg.GET("manufacturers/", h.handleManufacturersList)
	rg.GET("manufacturers/:manufacturer_id/", h.handleManufacturersEdit)
	rg.POST("manufacturers/:manufacturer_id/delete/", h.handleManufacturersDelete)
	// rg.GET("products/", handleProductsList)
	// rg.GET("products/:product_id/", handleProductsEdit)
}

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
	c.Redirect(http.StatusSeeOther, PrefixAdmin+adminRedirectLocation)
}
