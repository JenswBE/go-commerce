package handler

import (
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entity"
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
		c.String(400, "Parameter "+name+" is mandatory")
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

// errToResponse checks if the provided error is a GoComError.
// If yes, status and embedded error message are returned.
// If no, status is 500 and provided error message are returned.
func errToResponse(e error) (int, string) {
	if err, ok := e.(*entity.GoComError); ok {
		return err.Status, err.Err.Error()
	}
	return 500, e.Error()
}
