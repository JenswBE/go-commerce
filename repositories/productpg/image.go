package productpg

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories/productpg/internal"
)

func (r *ProductPostgres) GetImage(id entities.ID) (*entities.Image, error) {
	image := &internal.Image{}
	err := r.db.Take(image, "id = ?", id.String()).Error
	if err != nil {
		return nil, translatePgError(err, image, id.String())
	}
	return image.ToEntity(), nil
}

func (r *ProductPostgres) UpdateImage(id entities.ID, ownerID entities.ID, newOrder int) ([]*entities.Image, error) {
	// Fetch image
	image := &internal.Image{}
	err := r.db.Take(image, "id = ? AND owner_id = ?", id.String(), ownerID.String()).Error
	if err != nil {
		return nil, translatePgError(err, image, id.String())
	}

	// Order is the same => Ignore update request
	if image.Order == newOrder {
		return []*entities.Image{image.ToEntity()}, nil
	}

	// Fetch image with same new order (if any)
	imageSameOrder := &internal.Image{}
	err = r.db.Take(imageSameOrder, `owner_id = ? AND "order" = ?`, ownerID.String(), newOrder).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, translatePgError(err, imageSameOrder, "")
	}

	// Image with same new order exists => Update with old order
	if err == nil {
		imageSameOrder.Order = image.Order
		err = r.db.Save(imageSameOrder).Error
		if err != nil {
			return nil, translatePgError(err, imageSameOrder, imageSameOrder.ID)
		}
	}

	// Update image
	image.Order = newOrder
	err = r.db.Save(image).Error
	if err != nil {
		return nil, translatePgError(err, image, image.ID)
	}

	// Update successful
	if imageSameOrder.ID == "" {
		// No order swapped
		return []*entities.Image{image.ToEntity()}, nil
	}
	// Order swapped
	images := []*entities.Image{
		image.ToEntity(),
		imageSameOrder.ToEntity(),
	}
	return images, nil
}

func (r *ProductPostgres) DeleteImage(id entities.ID) error {
	// Delete image
	var images []internal.Image
	err := r.db.
		Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Delete(&images, "id = ?", id.String()).
		Error
	if err != nil {
		return translatePgError(err, &internal.Image{}, id.String())
	}

	// Return error if not found
	if len(images) == 0 {
		return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_IMAGE, id.String(), err)
	}
	return nil
}
