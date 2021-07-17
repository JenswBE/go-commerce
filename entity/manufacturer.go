package entity

import "net/url"

// Manufacturer data
type Manufacturer struct {
	ID         ID
	Name       string
	WebsiteURL string
	Image      *Image
}

// Validate validates the manufacturer data
func (m *Manufacturer) Validate() error {
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
