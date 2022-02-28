package image

import (
	"github.com/MijPeter/saxa/internal/db"
	"github.com/MijPeter/saxa/internal/error"
	"github.com/MijPeter/saxa/internal/model"
)

func Fetch(name string) (*model.Image, error) {
	image, _ := db.GetImage(name)

	if image == nil {
		return nil, httperror.IMAGE_NOT_FOUND
	}

	return image, nil
}

func Create(name string, content []byte) (*model.Image, error) {
	if imageDb, _ := db.GetImage(name); imageDb != nil {
		return nil, httperror.IMAGE_NAME_EXISTS
	}

	image := &model.Image{Name: name, Content: content}
	return db.SaveImage(name, image)
}

func Update(name string, newName string, content []byte) (*model.Image, error) {
	image, _ := db.GetImage(name)

	if image == nil {
		return nil, httperror.IMAGE_NOT_FOUND
	}

	image.Content = content
	return db.SaveImage(name, image)
}

// Returns json containing all names of stored images.
func Query() ([]string, error) {
	return db.GetNames(), nil
}
