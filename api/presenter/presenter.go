package presenter

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/sanitizer"
	"github.com/JenswBE/go-commerce/utils/shortid"
)

type Presenter struct {
	shortIDService   shortid.Service
	sanitizerService sanitizer.Service
}

func New(shortIDService shortid.Service, sanitizerService sanitizer.Service) *Presenter {
	return &Presenter{
		shortIDService:   shortIDService,
		sanitizerService: sanitizerService,
	}
}

func (p *Presenter) ParseID(id string) (entities.ID, error) {
	pID, err := p.shortIDService.Decode(id)
	if err != nil {
		// Parsing of short ID failed, try to parse as UUID
		var uuidErr error
		pID, uuidErr = entities.NewIDFromString(id)
		if uuidErr != nil {
			// UUID parsing failed => Return original error
			return entities.NewNilID(), entities.NewError(400, openapi.GOCOMERRORCODE_INVALID_ID, id, err)
		}
	}
	return pID, nil
}

func (p *Presenter) EncodeID(id entities.ID) string {
	return p.shortIDService.Encode(id)
}

func (p *Presenter) EncodeIDList(ids []entities.ID) []string {
	output := make([]string, 0, len(ids))
	for _, id := range ids {
		output = append(output, p.shortIDService.Encode(id))
	}
	return output
}

// Sanitizes input to a plain string.
// Shortcut for p.sanitizerService.String()
func (p *Presenter) String(input string) string {
	return p.sanitizerService.String(input)
}

// Sanitizes input to a very restricted HTML subset for content of type HTML
// Shortcut for p.sanitizerService.ContentHTML()
func (p *Presenter) ContentHTML(input string) string {
	return p.sanitizerService.ContentHTML(input)
}
