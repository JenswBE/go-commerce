package content

import (
	"github.com/JenswBE/go-commerce/entities"
)

// GetContent fetches a single content by ID
func (s *Service) GetContent(name string) (*entities.Content, error) {
	// Fetch content
	content, err := s.db.GetContent(name)
	if err != nil {
		return nil, err
	}

	// Get successful
	return content, nil
}

// ListContents fetches all contents
func (s *Service) ListContent() ([]*entities.Content, error) {
	// Fetch contents
	contents, err := s.db.ListContent()
	if err != nil {
		return nil, err
	}

	// Get successful
	return contents, nil
}

// CreateContent creates a new content
func (s *Service) CreateContent(content *entities.Content) (*entities.Content, error) {
	// Validate entity
	err := content.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.CreateContent(content)
}

// UpdateContent persists the provided content
func (s *Service) UpdateContent(content *entities.Content) (*entities.Content, error) {
	// Validate entity
	err := content.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.UpdateContent(content)
}

// DeleteContent deletes a single content by ID
func (s *Service) DeleteContent(name string) error {
	return s.db.DeleteContent(name)
}
