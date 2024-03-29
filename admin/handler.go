package admin

import (
	"errors"
	"io"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/JenswBE/go-commerce/admin/auth"
	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
	"github.com/JenswBE/go-commerce/config"
	baseEntities "github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/usecases/content"
	"github.com/JenswBE/go-commerce/usecases/product"
)

const (
	PrefixAdmin = "/admin/"
	IDNew       = "new"
)

const (
	pathLogin = "login/"
)

type Handler struct {
	features             config.Features
	contentService       content.Usecase
	productService       product.Usecase
	oidcClient           *auth.OIDCClient
	sessionAuthenticator *auth.SessionAuthenticator
	sessionAuthKey       [64]byte
	sessionEncKey        [32]byte
}

func NewHandler(features config.Features, productService product.Usecase, contentService content.Usecase, oidcClient *auth.OIDCClient, sessionAuthKey [64]byte, sessionEncKey [32]byte) *Handler {
	handler := &Handler{
		features:       features,
		productService: productService,
		contentService: contentService,
		oidcClient:     oidcClient,
		sessionAuthKey: sessionAuthKey,
		sessionEncKey:  sessionEncKey,
	}
	if oidcClient != nil {
		handler.sessionAuthenticator = auth.NewSessionAuthenticator(time.Hour * 24 * 7)
	}
	return handler
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	// Register middlewares
	notAuthenticatedGroup := r.Group(PrefixAdmin)
	cookieStore := cookie.NewStore(h.sessionAuthKey[:], h.sessionEncKey[:])
	cookieStore.Options(sessions.Options{
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
		Path:     "/",
	})
	notAuthenticatedGroup.Use(sessions.Sessions("gocom", cookieStore))
	rg := notAuthenticatedGroup.Group("")
	if h.sessionAuthenticator != nil {
		rg.Use(h.sessionAuthenticator.MW(PrefixAdmin + pathLogin))
	}

	// Register static routes
	staticFS, err := fs.Sub(htmlContent, "html/static")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to define sub FS for static content")
	}
	notAuthenticatedGroup.StaticFS("static", http.FS(staticFS))

	// Register dynamic routes
	rg.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, strings.ToLower(string(h.features.StartpageFeature))+"/")
	})
	notAuthenticatedGroup.GET(pathLogin, h.handleLogin)
	notAuthenticatedGroup.POST(pathLogin, h.handleLogin)
	notAuthenticatedGroup.GET(pathLogin+"oidc_redirect/", h.handleLoginOIDCRedirect)
	rg.GET("logout/", h.handleLogout)
	if h.features.Categories.Enabled {
		rg.GET("categories/", h.handleCategoriesList)
		rg.GET("categories/:category_id/", h.handleCategoriesFormGET)
		rg.POST("categories/:category_id/", h.handleCategoriesFormPOST)
		rg.POST("categories/:category_id/update_order/", h.handleCategoriesUpdateOrder)
		rg.POST("categories/:category_id/delete/", h.handleCategoriesDelete)
	}
	if h.features.Content.Enabled {
		rg.GET("content/", h.handleContentList)
		rg.GET("content/:content_name/", h.handleContentFormGET)
		rg.POST("content/:content_name/", h.handleContentFormPOST)
		rg.POST("content/:content_name/clear", h.handleContentClear)
	}
	if h.features.Events.Enabled {
		rg.GET("events/", h.handleEventsList)
		rg.GET("events/:event_id/", h.handleEventsFormGET)
		rg.POST("events/:event_id/", h.handleEventsFormPOST)
		rg.POST("events/:event_id/delete/", h.handleEventsDelete)
	}
	if h.features.Manufacturers.Enabled {
		rg.GET("manufacturers/", h.handleManufacturersList)
		rg.GET("manufacturers/:manufacturer_id/", h.handleManufacturersFormGET)
		rg.POST("manufacturers/:manufacturer_id/", h.handleManufacturersFormPOST)
		rg.POST("manufacturers/:manufacturer_id/delete/", h.handleManufacturersDelete)
		rg.GET("manufacturers/:manufacturer_id/image/", h.handleManufacturersImageGET)
		rg.POST("manufacturers/:manufacturer_id/image/", h.handleManufacturersImagePOST)
		rg.POST("manufacturers/:manufacturer_id/image/:image_id/delete/", h.handleManufacturersImageDelete)
	}
	if h.features.Products.Enabled {
		rg.GET("products/", h.handleProductsList)
		rg.GET("products/:product_id/", h.handleProductsFormGET)
		rg.POST("products/:product_id/", h.handleProductsFormPOST)
		rg.POST("products/:product_id/delete/", h.handleProductsDelete)
		rg.GET("products/:product_id/images/", h.handleProductsImagesGET)
		rg.POST("products/:product_id/images/", h.handleProductsImagesPOST)
		rg.POST("products/:product_id/images/:image_id/update_order/", h.handleProductsImagesUpdateOrder)
		rg.POST("products/:product_id/images/:image_id/delete/", h.handleProductsImagesDelete)
	}
	if h.features.Services.Enabled {
		rg.GET("service_categories/", h.handleServiceCategoriesList)
		rg.GET("service_categories/:service_category_id/", h.handleServiceCategoriesFormGET)
		rg.POST("service_categories/:service_category_id/", h.handleServiceCategoriesFormPOST)
		rg.POST("service_categories/:service_category_id/update_order/", h.handleServiceCategoriesUpdateOrder)
		rg.POST("service_categories/:service_category_id/delete/", h.handleServiceCategoriesDelete)
		rg.GET("service_categories/:service_category_id/services/", h.handleServicesList)
		rg.GET("service_categories/:service_category_id/services/:service_id/", h.handleServicesFormGET)
		rg.POST("service_categories/:service_category_id/services/:service_id/", h.handleServicesFormPOST)
		rg.POST("service_categories/:service_category_id/services/:service_id/update_order/", h.handleServicesUpdateOrder)
		rg.POST("service_categories/:service_category_id/services/:service_id/delete/", h.handleServicesDelete)
	}
}

func parseID(uuidInput string, objectType i18n.ObjectType) (baseEntities.ID, error) {
	output, err := baseEntities.NewIDFromString(uuidInput)
	if err != nil {
		return baseEntities.NewNilID(), errors.New(i18n.InvalidUUID(objectType, uuidInput))
	}
	return output, nil
}

func html(c *gin.Context, code int, template entities.Template) {
	c.HTML(code, template.GetTemplateName(), template)
}

func htmlWithFlashes(c *gin.Context, template entities.Template) {
	// Get and convert flashes
	session := sessions.Default(c)
	flashes := session.Flashes()
	messages := make([]entities.Message, 0, len(flashes))
	for _, flash := range flashes {
		flashString, ok := flash.(string)
		if !ok {
			c.String(http.StatusInternalServerError, "Unable to cast flash value to string. Flash type is %T. Flashes: %v", flash, flashes)
		}
		messages = append(messages, entities.ParseMessage(flashString))
	}
	template.SetMessages(messages)

	// Save session
	err := session.Save()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to save session after retrieving flashes: %v. Flashes: %v", err, flashes)
		return
	}

	// Display page
	c.HTML(http.StatusOK, template.GetTemplateName(), template)
}

func redirect(c *gin.Context, adminRedirectLocation string) {
	c.Redirect(http.StatusSeeOther, PrefixAdmin+adminRedirectLocation)
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
			}
			return nil, err
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
