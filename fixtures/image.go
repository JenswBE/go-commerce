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
		URLs:      map[string]string{ImageConfigString: "http://image.test"},
	}
}

func ImageSlice() []*entities.Image {
	return []*entities.Image{
		Image(),
	}
}

func ImageConfig() imageproxy.ImageConfig {
	return imageproxy.ImageConfig{
		Width:        300,
		Height:       200,
		ResizingType: imageproxy.ResizingTypeFill,
	}
}

func ImageConfigMap() map[string]imageproxy.ImageConfig {
	return map[string]imageproxy.ImageConfig{
		ImageConfigString: ImageConfig(),
	}
}

const ImageConfigString = "300_200_FILL"
const ImageConfigQuery = "img=" + ImageConfigString

// #############################
// #          OPENAPI          #
// #############################

func ImageOpenAPI() *openapi.Image {
	return &openapi.Image{
		Id:    ImageID,
		Ext:   ".jpg",
		Urls:  map[string]string{ImageConfigString: "http://image.test"},
		Order: 1,
	}
}

func ImageOpenAPISlice() []openapi.Image {
	return []openapi.Image{
		*ImageOpenAPI(),
	}
}
