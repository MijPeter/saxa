package util

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"

	httperror "github.com/MijPeter/saxa/internal/error"
)

type file struct {
	name string
	header []*multipart.FileHeader
}

func ParseFile(r *http.Request) (string, []byte, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return "", nil, err
	}

	files := r.MultipartForm.File
	if len(files) != 1 {
		return "", nil, httperror.INCORRECT_FILE_NUMBER
	}
	
	fileSlice := make([]file, 0, 1) 
	 
	for name, headers := range files {
		fileSlice = append(fileSlice, file{name, headers})
	}

	if len(fileSlice[0].header) != 1 {
		return "", nil, httperror.INCORRECT_FILE_HEADER
	}

	header := fileSlice[0].header[0]
	
	physicalFile, err := header.Open()
	if err != nil {
		return "", nil, err
	}

	defer physicalFile.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, physicalFile); err != nil {
		return "", nil, err
	}

	return fileSlice[0].name, buf.Bytes(), nil
}