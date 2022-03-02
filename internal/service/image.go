package image

import (
	"github.com/MijPeter/saxa/internal/error"
	"github.com/MijPeter/saxa/internal/model"
	"github.com/MijPeter/saxa/internal/persistence"
)

func Fetch(name string) (*model.Image, error) {
	image, err := persistence.GetImage(name)

	if err != nil {
		return nil, err
	}

	if image == nil {
		return nil, httperror.IMAGE_NOT_FOUND
	}

	return image, nil
}

func Create(name string, content []byte) (*model.Image, error) {
	image := &model.Image{Name: name, Content: content}
	return image, persistence.SaveImage(name, image)
}

// Returns json containing all names of stored images.
func Query() ([]string, error) {
	return persistence.GetNames()
}
