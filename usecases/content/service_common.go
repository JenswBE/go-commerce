package content

import "github.com/JenswBE/go-commerce/entities"

var _ Usecase = &Service{}

type Service struct {
	db DatabaseRepository
}

func NewService(db DatabaseRepository, content []entities.Content) (*Service, error) {
	// Build service
	service := &Service{db: db}

	// Ensure content is present in DB
	err := service.alignContentWithDB(content)
	if err != nil {
		return nil, err
	}

	// Build and return service
	return service, nil
}

func (s *Service) alignContentWithDB(contentList []entities.Content) error {
	// Fetch current content
	currentContentList, err := s.ListContent()
	if err != nil {
		return err
	}

	// Process content
	for _, newContent := range contentList {
		currentContent := findContent(currentContentList, newContent.Name)
		if currentContent == nil {
			// Content not found => Create
			_, err = s.CreateContent(&newContent)

		} else {
			// Content type is same => Skip
			if currentContent.ContentType == newContent.ContentType {
				continue
			}

			// Update content type
			currentContent.ContentType = newContent.ContentType
			_, err = s.UpdateContent(currentContent)
		}
		if err != nil {
			return err
		}
	}

	// Content successful aligned
	return nil
}

func findContent(list []*entities.Content, name string) *entities.Content {
	for _, content := range list {
		if content.Name == name {
			return content
		}
	}
	return nil
}
