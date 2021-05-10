package handler

import (
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// parseParamID tries to parse parameter "id" as an UUID or short ID.
// On failure, an error is set on the Gin context.
//
//   p := presenter.New()
//   id, ok := parseParamID(c, p)
//   if !ok {
// 	   return // Response already set on Gin context
//   }
func parseParamID(c *gin.Context, p *presenter.Presenter) (uuid.UUID, bool) {
	// Parse param
	pID, ok := c.Params.Get("id")
	if !ok {
		c.String(400, "Parameter id is mandatory")
		return uuid.Nil, false
	}

	// Parse ID
	id, err := p.ParseID(pID)
	if err != nil {
		c.String(400, err.Error())
		return uuid.Nil, false
	}

	// Parse successful
	return id, true
}
