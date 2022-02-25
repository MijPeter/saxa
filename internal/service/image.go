package image

import (
	"errors"

	"github.com/MijPeter/saxa/internal/db"
	"github.com/MijPeter/saxa/internal/error"
	"github.com/MijPeter/saxa/internal/model"
)

func Fetch(name string) (*model.Image, err.HttpError) {
	image, _ := db.GetImage(name)

	if image == nil {
		return nil, err.IMAGE_NOT_FOUND
	}

	errors.New()
	
	return image, nil
}

func Create(name string, content []byte) (*model.Image, err.HttpError) {
	if imageDb, _ := db.GetImage(name); imageDb != nil {
		return nil, err.IMAGE_NAME_EXISTS
	}

	image := &model.Image{Name: name, Content: content}
	return db.SaveImage(name, image)
}

func Update(name string, newName string, content []byte) (*model.Image, err.HttpError) {
	image, _ := db.GetImage(name)

	if image == nil {
		return nil, err.IMAGE_NOT_FOUND
	}

	image.Content = content
	return db.SaveImage(name, image)
}

// Returns json containing all names of stored images.
func Query() ([]string, error) {
	return db.GetNames(), nil
}
