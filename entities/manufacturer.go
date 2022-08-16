package entities

import (
	"net/url"
	"strings"
)

// Manufacturer data
type Manufacturer struct {
	ID         ID
	Name       string
	WebsiteURL string
	Image      *Image
}

func (m *Manufacturer) Clean() {
	m.Name = strings.TrimSpace(m.Name)
	m.WebsiteURL = strings.TrimSpace(m.WebsiteURL)
}

// Validate cleans and validates the manufacturer data
func (m *Manufacturer) Validate() error {
	// Clean entity
	m.Clean()

	// Validate simple fields
	if m.Name == "" {
		return ErrInvalidEntity
	}

	// Validate URL
	if m.WebsiteURL != "" {
		_, err := url.Parse(m.WebsiteURL)
		if err != nil {
			return ErrInvalidEntity
		}
	}

	// Entity is valid
	return nil
}
