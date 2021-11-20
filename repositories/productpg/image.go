package productpg

import (
	"errors"

	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories/productpg/internal"
	"gorm.io/gorm"
)

func (r *ProductPostgres) GetImage(id entities.ID) (*entities.Image, error) {
	image := &internal.Image{}
	err := r.db.Take(image, "id = ?", id).Error
	if err != nil {
		return nil, translatePgError(err, image, id.String())
	}
	return image.ToEntity(), nil
}

func (r *ProductPostgres) UpdateImage(id entities.ID, ownerID entities.ID, newOrder int) ([]*entities.Image, error) {
	// Fetch image
	image := &internal.Image{}
	err := r.db.Take(image, "id = ? AND owner_id = ?", id, ownerID).Error
	if err != nil {
		return nil, translatePgError(err, image, id.String())
	}

	// Order is the same => Ignore update request
	if image.Order == newOrder {
		return []*entities.Image{image.ToEntity()}, nil
	}

	// Fetch image with same new order (if any)
	imageSameOrder := &internal.Image{}
	err = r.db.Take(imageSameOrder, `owner_id = ? AND "order" = ?`, ownerID, newOrder).Error
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
	} else {
		// Order swapped
		images := []*entities.Image{
			image.ToEntity(),
			imageSameOrder.ToEntity(),
		}
		return images, nil
	}
}

func (r *ProductPostgres) DeleteImage(id entities.ID) error {
	err := r.db.Delete(&internal.Image{}, "id = ?", id).Error
	if err != nil {
		return translatePgError(err, &internal.Image{}, id.String())
	}
	return nil
}
