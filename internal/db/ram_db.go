package db

import (
	"github.com/MijPeter/saxa/internal/model"
)

var images = make(map[string]*model.Image)

func GetImage(name string) (*model.Image, error) {
	return images[name], nil
}

func SaveImage(name string, image *model.Image) (*model.Image, error) {
	images[name] = image
	return image, nil
}

func GetNames() []string {
	names := make([]string, 0, len(images))
	for _, image := range images {
		names = append(names, image.Name)
	}
	return names
}
