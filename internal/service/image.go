package image

import "github.com/MijPeter/saxa/internal/model"

func Fetch(name string) model.Image {
	return model.Image{Name: "", Content: nil}
}

func Create(name string, content []byte) model.Image {
	return model.Image{Name: "", Content: nil}
}

func Update(name string, newName string, content []byte) model.Image {
	return model.Image{Name: "", Content: nil}
}

// Returns json containing all stored images.
func Query() []model.Image {
	return []model.Image{}
}
