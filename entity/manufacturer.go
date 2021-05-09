package entity

import "net/url"

// Manufacturer data
type Manufacturer struct {
	ID         ID
	Name       string
	WebsiteURL string
}

// NewManufacturer creates a new manufacturer
func NewManufacturer(name, websiteURL string) (*Manufacturer, error) {
	b := &Manufacturer{
		ID:         NewID(),
		Name:       name,
		WebsiteURL: websiteURL,
	}
	err := b.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return b, nil
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
