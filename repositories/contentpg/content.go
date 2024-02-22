package contentpg

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories"
	"github.com/JenswBE/go-commerce/repositories/contentpg/internal"
)

func (r *ContentPostgres) GetContent(name string) (*entities.Content, error) {
	content := &internal.Content{}
	err := r.db.Take(content, "name = ?", name).Error
	if err != nil {
		return nil, translatePgError(err, content, name)
	}
	return content.ToEntity(), nil
}

func (r *ContentPostgres) ListContent() ([]*entities.Content, error) {
	contents := []*internal.Content{}
	err := r.db.Order("LOWER(name)").Find(&contents).Error
	if err != nil {
		return nil, translatePgError(err, contents, "")
	}
	return repositories.ToEntitiesList(contents, (*internal.Content).ToEntity), nil
}

func (r *ContentPostgres) CreateContent(e *entities.Content) (*entities.Content, error) {
	content := internal.ContentEntityToPg(e)
	err := r.db.Create(content).Error
	if err != nil {
		return nil, translatePgError(err, content, content.Name)
	}
	return content.ToEntity(), nil
}

func (r *ContentPostgres) UpdateContent(e *entities.Content) (*entities.Content, error) {
	content := internal.ContentEntityToPg(e)
	err := r.db.Save(content).Error
	if err != nil {
		return nil, translatePgError(err, content, content.Name)
	}
	return content.ToEntity(), nil
}

func (r *ContentPostgres) DeleteContent(name string) error {
	err := r.db.Delete(&internal.Content{}, "name = ?", name).Error
	if err != nil {
		return translatePgError(err, &internal.Content{}, name)
	}
	return nil
}
