package handler

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// parseIDParam tries to parse parameter with the given name as an UUID or short ID.
// On failure, an error is set on the Gin context.
//
//   p := presenter.New()
//   id, ok := parseIDParam(c, "id", p)
//   if !ok {
// 	   return // Response already set on Gin context
//   }
func parseIDParam(c *gin.Context, name string, p *presenter.Presenter) (uuid.UUID, bool) {
	// Parse param
	pID, ok := c.Params.Get(name)
	if !ok {
		err := entities.NewError(400, openapi.ERRORCODE_PARAMETER_MISSING, name, nil)
		c.JSON(errToResponse(err))
		return uuid.Nil, false
	}

	// Parse ID
	id, err := p.ParseID(pID)
	if err != nil {
		c.JSON(errToResponse(err))
		return uuid.Nil, false
	}

	// Parse successful
	return id, true
}

// errToResponse checks if the provided error is a GoComError.
// If yes, status and embedded error message are returned.
// If no, status is 500 and provided error message are returned.
func errToResponse(e error) (int, *entities.GoComError) {
	if err, ok := e.(*entities.GoComError); ok {
		return err.Status, err
	}
	return 500, entities.NewError(500, openapi.ERRORCODE_UNKNOWN_ERROR, "", e).(*entities.GoComError)
}
