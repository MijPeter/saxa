package persistence

import (
	"database/sql"

	"github.com/MijPeter/saxa/db"
	"github.com/MijPeter/saxa/internal/model"
)


func GetImage(name string) (*model.Image, error) {
	var content []byte
	err := db.DB.QueryRow("select * from image where name=$1", name).Scan(&name, &content)
	
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &model.Image{Name: name, Content: content}, nil
}

func SaveImage(name string, image *model.Image) error {
	_, err := db.DB.Exec("insert into image (name, content) values ($1, $2)", name, image.Content)
	
	return err
}

func GetNames() ([]string, error) {
	rows, err := db.DB.Query("select name from image")
	if err == sql.ErrNoRows {
		return []string{}, err
	}


	if err != nil {
		return []string{}, err
	}

	var names []string = []string{}

	for rows.Next() {
		var name string
		rows.Scan(&name)
		names = append(names, name)
	}

	return names, nil
}