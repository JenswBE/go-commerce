package admin

import (
	"errors"
	"net/http"
	"time"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
	"github.com/JenswBE/go-commerce/usecases/product"
	"github.com/JenswBE/go-commerce/utils/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

const PrefixAdmin = "/admin/"

const (
	objectTypeManufacturer = "merk"
	pathLogin              = "login/"
)

type AdminHandler struct {
	productService product.Usecase
	authVerifier   auth.Verifier
	jwtSigningKey  [64]byte
	sessionAuthKey [64]byte
}

func NewAdminGUIHandler(productService product.Usecase, sessionAuthKey [64]byte, authVerifier auth.Verifier, jwtSigningKey [64]byte) (*AdminHandler, error) {
	return &AdminHandler{
		productService: productService,
		authVerifier:   authVerifier,
		jwtSigningKey:  jwtSigningKey,
		sessionAuthKey: sessionAuthKey,
	}, nil
}

func (h *AdminHandler) RegisterRoutes(r *gin.Engine) {
	// Register middlewares
	notAuthenticatedGroup := r.Group(PrefixAdmin)
	notAuthenticatedGroup.Use(sessions.Sessions("gocom", cookie.NewStore(h.sessionAuthKey[:])))
	rg := notAuthenticatedGroup.Group("")
	rg.Use(validateAndRefreshToken(h.jwtSigningKey))

	// Register static routes
	rg.Static("static", "admin/html/static")

	// Register dynamic routes
	rg.GET("/", func(c *gin.Context) { c.Redirect(http.StatusTemporaryRedirect, "products/") })
	// rg.GET("categories/", handleCategoriesList)
	notAuthenticatedGroup.GET(pathLogin, h.handleLogin)
	notAuthenticatedGroup.POST(pathLogin, h.handleLogin)
	rg.GET("logout/", h.handleLogout)
	rg.GET("manufacturers/", h.handleManufacturersList)
	rg.GET("manufacturers/:manufacturer_id/", h.handleManufacturersEdit)
	rg.GET("manufacturers/:manufacturer_id/delete/", h.handleManufacturersDelete)
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

func validateAndRefreshToken(jwtSigningKey [64]byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from session
		session := sessions.Default(c)
		rawToken := session.Get("token")
		if rawToken == nil {
			c.Redirect(http.StatusSeeOther, PrefixAdmin+pathLogin)
			c.Abort()
			return
		}

		// Validate type of token
		tokenString, ok := rawToken.(string)
		if !ok {
			session.Clear()
			if err := session.Save(); err != nil {
				log.Warn().Err(err).Interface("token", rawToken).Msg("Failed to clear session")
				redirectWithMessage(c, session, entities.MessageTypeError, "Opkuisen van sessie mislukt. Probeer opnieuw in te loggen.", pathLogin)
			} else {
				redirectWithMessage(c, session, entities.MessageTypeError, "Ongeldig type token in sessie. Probeer opnieuw in te loggen.", pathLogin)
			}
			c.Abort()
			return
		}

		// Validate token
		tokenClaims, err := auth.ValidateJWT(tokenString, jwtSigningKey)
		if err != nil {
			c.Redirect(http.StatusSeeOther, PrefixAdmin+pathLogin)
			c.Abort()
			return
		}

		// Refresh token if older than 1 day
		if tokenClaims.IssuedAt != nil && tokenClaims.IssuedAt.Time.Before(time.Now().AddDate(0, 0, 1)) {
			err = setNewTokenInSession(c, jwtSigningKey)
			if err != nil {
				log.Warn().Err(err).Msg("Failed to refresh token in session")
				redirectWithMessage(c, session, entities.MessageTypeError, "Verversen van token mislukt. Probeer opnieuw in te loggen.", pathLogin)
				c.Abort()
				return
			}
		}
	}
}
