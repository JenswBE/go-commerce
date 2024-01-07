package entities

import (
	"strings"
	"time"

	"github.com/JenswBE/go-commerce/api/openapi"
)

// Service data
type Service struct {
	ID          ID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	// Price of the service in cents (1/100)
	Price AmountInCents
	// Order (priority) of within the service category.
	// 1 = highest, inf = lowest
	Order             int
	ServiceCategoryID ID
}

func (s *Service) Clean() {
	s.Name = strings.TrimSpace(s.Name)
	s.Description = strings.TrimSpace(s.Description)
}

// Validate cleans and validates the service data
func (s *Service) Validate() error {
	// Clean entity
	s.Clean()

	// Validate simple fields
	if s.Name == "" {
		return NewError(400, openapi.GOCOMERRORCODE_SERVICE_NAME_EMPTY, s.ID.String(), nil)
	}
	if s.Price.Int() < 0 {
		return NewError(400, openapi.GOCOMERRORCODE_SERVICE_PRICE_NEGATIVE, s.ID.String(), nil)
	}
	if s.Order < 0 {
		return NewError(400, openapi.GOCOMERRORCODE_SERVICE_ORDER_NEGATIVE, s.ID.String(), nil)
	}

	// Entity is valid
	return nil
}
