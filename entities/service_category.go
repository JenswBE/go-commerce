package entities

import (
	"strings"

	"github.com/JenswBE/go-commerce/api/openapi"
)

type ServiceCategory struct {
	ID   ID
	Name string
	// Order (priority) of the category.
	// 1 = highest, inf = lowest
	Order      int
	ServiceIDs []ID
}

func (c *ServiceCategory) Clean() {
	c.Name = strings.TrimSpace(c.Name)
}

// Validate cleans and validates the category data
func (c *ServiceCategory) Validate() error {
	// Clean entity
	c.Clean()

	// Validate simple fields
	if c.Name == "" {
		return NewError(400, openapi.GOCOMERRORCODE_SERVICE_CATEGORY_NAME_EMPTY, c.ID.String(), nil)
	}
	if c.Order < 0 {
		return NewError(400, openapi.GOCOMERRORCODE_SERVICE_CATEGORY_ORDER_NEGATIVE, c.ID.String(), nil)
	}

	// Entity is valid
	return nil
}

// ResolvedServiceCategory is a service category for which related entities
// are included. This way all information is immediately at hand.
type ResolvedServiceCategory struct {
	ServiceCategory
	Services []*Service
}
