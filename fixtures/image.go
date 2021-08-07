package fixtures

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/google/uuid"
)

// #############################
// #           ENTITY          #
// #############################

func Image() *entities.Image {
	return &entities.Image{
		ID:        uuid.MustParse(ImageID),
		Extension: ".jpg",
		Order:     1,
		URL:       "http://image.test",
	}
}

func ImageSlice() []*entities.Image {
	return []*entities.Image{
		Image(),
	}
}

func ImageConfig() *imageproxy.ImageConfig {
	return &imageproxy.ImageConfig{
		Width:        300,
		Height:       200,
		ResizingType: imageproxy.ResizingTypeFill,
	}
}

const ImageConfigQuery = "img_w=300&img_h=200&img_r=FILL"

// #############################
// #          OPENAPI          #
// #############################

func ImageOpenAPI() *openapi.Image {
	return &openapi.Image{
		Id:    ImageID,
		Ext:   ".jpg",
		Url:   "http://image.test",
		Order: 1,
	}
}

func ImageOpenAPISlice() []openapi.Image {
	return []openapi.Image{
		*ImageOpenAPI(),
	}
}

func ImageListOpenAPI() *openapi.ImageList {
	return openapi.NewImageList(ImageOpenAPISlice())
}
