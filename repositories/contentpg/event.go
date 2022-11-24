package contentpg

import (
	"gorm.io/gorm/clause"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories/contentpg/internal"
)

func (r *ContentPostgres) GetEvent(id entities.ID) (*entities.Event, error) {
	event := &internal.Event{}
	err := r.db.Take(event, "id = ?", id.String()).Error
	if err != nil {
		return nil, translatePgError(err, event, id.String())
	}
	return event.ToEntity(), nil
}

func (r *ContentPostgres) ListEvents(includePastEvents bool) ([]*entities.Event, error) {
	events := []*internal.Event{}
	query := r.db.Order(`"start", "end", LOWER("name")`)
	if !includePastEvents {
		query = query.Where(`"end" > now()`)
	}
	err := query.Find(&events).Error
	if err != nil {
		return nil, translatePgError(err, events, "")
	}
	return internal.EventsListPgToEntity(events), nil
}

func (r *ContentPostgres) CreateEvent(e *entities.Event) (*entities.Event, error) {
	m := internal.EventEntityToPg(e)
	err := r.db.Create(m).Error
	if err != nil {
		return nil, translatePgError(err, m, m.ID)
	}
	return m.ToEntity(), nil
}

func (r *ContentPostgres) UpdateEvent(e *entities.Event) (*entities.Event, error) {
	m := internal.EventEntityToPg(e)
	err := r.db.Save(m).Error
	if err != nil {
		return nil, translatePgError(err, m, e.ID.String())
	}
	return m.ToEntity(), nil
}

func (r *ContentPostgres) DeleteEvent(id entities.ID) error {
	// Delete event
	var events []internal.Event
	err := r.db.
		Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Delete(&events, "id = ?", id.String()).
		Error
	if err != nil {
		return translatePgError(err, &internal.Event{}, id.String())
	}

	// Return error if not found
	if len(events) == 0 {
		return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_EVENT, id.String(), err)
	}
	return nil
}
