package productpg

import (
	"github.com/JenswBE/go-commerce/entity"
	"github.com/JenswBE/go-commerce/repository/productpg/internal"
)

func (r *ProductPostgres) GetImage(id entity.ID) (*entity.Image, error) {
	image := &internal.Image{}
	err := r.db.Take(image, "id = ?", id).Error
	if err != nil {
		return nil, translatePgError(err, "image")
	}
	return internal.ImagePgToEntity(image), nil
}

func (r *ProductPostgres) DeleteImage(id entity.ID) error {
	err := r.db.Delete(&internal.Image{}, "id = ?", id).Error
	if err != nil {
		return translatePgError(err, "image")
	}
	return nil
}
